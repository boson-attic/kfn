package rust

import (
	"context"
	"fmt"
	"github.com/containers/image/types"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/image"
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"github.com/slinkydeveloper/kfn/pkg/util"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

const (
	rustRuntimeRemoteZip = "https://github.com/openshift-cloud-functions/faas-rust-runtime/archive/master.zip"
	buildEnvVariables    = "build-env"
	buildDevProfile      = "build-dev"
)

type rustLanguageManager struct {
	resourceLoader util.ResourceLoader
}

func (r rustLanguageManager) Bootstrap(functionName string, targetDirectory string) error {
	err := util.MkdirpIfNotExists(targetDirectory)
	if err != nil {
		return err
	}

	main, err := r.resourceLoader.LoadResource("function.rs")
	if err != nil {
		return err
	}

	return util.WriteFiles(
		targetDirectory,
		util.WriteDest{fmt.Sprintf("%s.rs", functionName), main},
	)
}

func (r rustLanguageManager) DownloadRuntimeIfRequired() error {
	if !util.FsExist(runtimeDirectory()) {
		if err := util.MkdirpIfNotExists(runtimeDirectory()); err != nil {
			return err
		}

		tempDir, err := ioutil.TempDir("", "faas-rust-runtime")
		if err != nil {
			return err
		}

		runtimeZip := path.Join(tempDir, "master.zip")

		if err := util.DownloadFile(rustRuntimeRemoteZip, runtimeZip); err != nil {
			return err
		}

		logrus.Infof("Downloading runtime from %s to %s", rustRuntimeRemoteZip, runtimeZip)

		if _, err := util.Unzip(runtimeZip, tempDir); err != nil {
			return err
		}

		logrus.Infof("Runtime unzipped to %s", tempDir)

		if err := util.Copy(path.Join(tempDir, "faas-rust-runtime-master", "src"), path.Join(runtimeDirectory(), "src")); err != nil {
			return err
		}

		if err := util.Copy(path.Join(tempDir, "faas-rust-runtime-master", "Cargo.toml"), path.Join(runtimeDirectory())); err != nil {
			return err
		}

		logrus.Infof("Copied runtime to %s", runtimeDirectory())
	} else {
		logrus.Infof("Using runtime cached in %s", runtimeDirectory())
	}
	return nil
}

func (r rustLanguageManager) CheckCompileDependencies() error {
	return util.CommandsExists("rustc", "cargo")
}

func (r rustLanguageManager) ConfigureEditingDirectory(mainFile string, functionConfiguration map[string][]string) (string, string, error) {
	functionFile := path.Join(config.EditingDir, "lib.rs")

	cargoToml, err := generateCargoToml(functionConfiguration)
	if err != nil {
		return "", "", err
	}

	err = util.WriteFiles(config.EditingDir, util.WriteDest{Filename: "Cargo.toml", Data: cargoToml})
	if err != nil {
		return "", "", err
	}

	err = util.Link(mainFile, functionFile)
	if err != nil {
		return "", "", err
	}

	return config.EditingDir, path.Join(config.EditingDir, "Cargo.toml"), nil
}

func (r rustLanguageManager) ConfigureTargetDirectory(mainFile string, functionConfiguration map[string][]string) error {
	if err := util.MkdirpIfNotExists(path.Join(config.TargetDir, "function")); err != nil {
		return err
	}

	err := util.Copy(mainFile, path.Join(config.TargetDir, "function", "lib.rs"))
	if err != nil {
		return err
	}

	cargoToml, err := generateCargoToml(functionConfiguration)
	if err != nil {
		return err
	}

	err = util.WriteFiles(path.Join(config.TargetDir, "function"), util.WriteDest{Filename: "Cargo.toml", Data: cargoToml})
	if err != nil {
		return err
	}

	return util.CopyContent(runtimeDirectory(), path.Join(config.TargetDir, "runtime"))
}

func (r rustLanguageManager) Compile(mainFile string, functionConfiguration map[string][]string) (string, []string, error) {
	devMode := false

	// Check profile for building
	dev, ok := functionConfiguration[buildDevProfile]
	if ok {
		var err error
		devMode, err = strconv.ParseBool(dev[0])
		if err != nil {
			devMode = false
		}
	}

	log.Printf("Using cargo dev profile: %v", devMode)

	var compileCommand *exec.Cmd
	if devMode {
		compileCommand = exec.Command("cargo", "build")
	} else {
		compileCommand = exec.Command("cargo", "build", "--release")
	}
	// Root Cargo.toml is in runtime dir in runtime
	compileCommand.Dir = path.Join(config.TargetDir, "runtime")
	// Configure proper logging
	compileCommand.Stdout = config.GetLoggerWriter()
	compileCommand.Stderr = config.GetLoggerWriter()
	compileCommand.Env = os.Environ()

	envFlags, ok := functionConfiguration[buildEnvVariables]
	if ok {
		for _, env := range envFlags {
			log.Printf("Adding env variable to cargo build: %s", env)
			compileCommand.Env = append(compileCommand.Env, env)
		}
	}

	err := compileCommand.Run()
	if err != nil {
		return "", nil, err
	}

	if devMode {
		return path.Join(config.TargetDir, "runtime", "target", "debug", "rust-faas"), nil, nil
	} else {
		return path.Join(config.TargetDir, "runtime", "target", "release", "rust-faas"), nil, nil
	}
}

func (r rustLanguageManager) BuildImage(systemContext *types.SystemContext, imageName string, imageTag string, mainExecutable string, additionalFiles []string) (image.FunctionImage, error) {
	builder, err := util.InitializeBuilder(context.TODO(), systemContext, "registry.access.redhat.com/ubi8/ubi")
	if err != nil {
		return image.FunctionImage{}, err
	}

	builder.SetPort("8080")

	err = util.Add(
		builder,
		util.BuildAdd{From: mainExecutable, To: "/usr/bin"},
	)
	if err != nil {
		return image.FunctionImage{}, err
	}

	builder.SetUser("1000")
	builder.SetCmd([]string{"/usr/bin/rust-faas"})
	builder.SetEntrypoint([]string{})

	return util.CommitImage(builder, systemContext, imageName, imageTag)
}

func NewRustLanguageManger() languages.LanguageManager {
	return rustLanguageManager{util.NewResourceLoader("../../templates/rust")}
}

func runtimeDirectory() string {
	return path.Join(config.RuntimeDir, "rust")
}

type CargoFile struct {
	Package      map[string]string `toml:"package"`
	Lib          map[string]string `toml:"lib"`
	Dependencies map[string]string `toml:"dependencies"`
}

func generateCargoToml(configurationEntries map[string][]string) ([]byte, error) {
	deps := make(map[string]string)
	deps["actix-web"] = "1.0.8"
	deps["serde_json"] = "1.0"
	deps["futures"] = "0.1.29"

	if depConfigs, ok := configurationEntries[util.DEPENDENCY]; ok {
		for _, dep := range depConfigs {
			splitted := strings.Split(dep, " ")
			if len(splitted) != 2 {
				return nil, fmt.Errorf("Invalid dependency entry: %v", dep)
			}
			deps[strings.Trim(splitted[0], " ")] = strings.Trim(splitted[1], " ")
		}
	}

	root := CargoFile{
		Package: map[string]string{
			"name":    "function",
			"version": "0.0.1",
			"edition": "2018",
		},
		Lib: map[string]string{
			"path": "lib.rs",
		},
		Dependencies: deps,
	}

	return toml.Marshal(root)

}
