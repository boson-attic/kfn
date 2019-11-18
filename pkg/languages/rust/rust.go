package rust

import (
	"context"
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"strings"

	"github.com/containers/image/types"
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/image"
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"github.com/slinkydeveloper/kfn/pkg/util"
)

const (
	rustRuntimeRemoteZip = "https://github.com/openshift-cloud-functions/faas-rust-runtime/archive/master.zip"
	buildEnvVariables    = "build-env"
	testEnvVariables     = "test-env"
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
	// musl-gcc is required for static linking the libc
	return util.CommandsExists("rustc", "cargo", "musl-gcc")
}

func (r rustLanguageManager) ConfigureEditingDirectory(mainFile string, functionConfiguration map[string][]string, editingDirectory string) (string, error) {
	functionFile := path.Join(editingDirectory, "lib.rs")

	cargoToml, err := generateCargoToml(functionConfiguration)
	if err != nil {
		return "", err
	}

	err = util.WriteFiles(editingDirectory, util.WriteDest{Filename: "Cargo.toml", Data: cargoToml})
	if err != nil {
		return "", err
	}

	err = util.Link(mainFile, functionFile)
	if err != nil {
		return "", err
	}

	// Configure test file
	testFile := util.UnitTestFile(mainFile)
	if util.FsExist(testFile) {
		err = util.MkdirpIfNotExists(path.Join(editingDirectory, "tests"))
		if err != nil {
			return "", err
		}

		err := util.Link(testFile, path.Join(editingDirectory, "tests", "lib_test.rs"))
		if err != nil {
			return "", err
		}
	}

	return editingDirectory, nil
}
func (r rustLanguageManager) ConfigureTargetDirectory(mainFile string, functionConfiguration map[string][]string, targetDirectory string) error {
	if err := util.MkdirpIfNotExists(path.Join(targetDirectory, "function")); err != nil {
		return err
	}

	err := util.Copy(mainFile, path.Join(targetDirectory, "function", "lib.rs"))
	if err != nil {
		return err
	}

	cargoToml, err := generateCargoToml(functionConfiguration)
	if err != nil {
		return err
	}

	err = util.WriteFiles(path.Join(targetDirectory, "function"), util.WriteDest{Filename: "Cargo.toml", Data: cargoToml})
	if err != nil {
		return err
	}

	// Configure test file
	testFile := util.UnitTestFile(mainFile)
	if util.FsExist(testFile) {
		err = util.MkdirpIfNotExists(path.Join(targetDirectory, "function", "tests"))
		if err != nil {
			return err
		}

		err := util.Copy(testFile, path.Join(targetDirectory, "function", "tests", "lib_test.rs"))
		if err != nil {
			return err
		}
	}

	return util.CopyContent(runtimeDirectory(), path.Join(targetDirectory, "runtime"))
}

func (r rustLanguageManager) UnitTest(mainFile string, functionConfiguration map[string][]string, targetDirectory string) error {
	err := util.RunCommand(
		"cargo",
		[]string{"test", "--target", "x86_64-unknown-linux-musl"},
		path.Join(targetDirectory, "function"),
		functionConfiguration[buildEnvVariables], functionConfiguration[testEnvVariables],
	)
	if err != nil {
		return errors.Wrap(err, "error occurred while testing function")
	}

	return nil
}

func (r rustLanguageManager) Compile(mainFile string, functionConfiguration map[string][]string, targetDirectory string) (string, []string, error) {
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

	log.Debugf("Using cargo dev profile: %v", devMode)

	var args []string
	if devMode {
		args = []string{"build", "--target", "x86_64-unknown-linux-musl"}
	} else {
		args = []string{"build", "--release", "--target", "x86_64-unknown-linux-musl"}
	}

	err := util.RunCommand("cargo", args, path.Join(targetDirectory, "runtime"), functionConfiguration[buildEnvVariables])
	if err != nil {
		return "", nil, errors.Wrap(err, "error occurred while trying to compile. Check if you installed correctly 'https://www.musl-libc.org/how.html' and musl rustc target with 'rustup target add x86_64-unknown-linux-musl'")
	}

	if devMode {
		return path.Join(targetDirectory, "runtime", "target", "x86_64-unknown-linux-musl", "debug", "rust-faas"), nil, nil
	} else {
		return path.Join(targetDirectory, "runtime", "target", "x86_64-unknown-linux-musl", "release", "rust-faas"), nil, nil
	}
}

func (r rustLanguageManager) BuildImage(systemContext *types.SystemContext, imageName string, imageTag string, mainExecutable string, additionalFiles []string, targetDirectory string) (image.FunctionImage, error) {
	builder, err := util.InitializeBuilder(context.TODO(), systemContext, "")
	if err != nil {
		return image.FunctionImage{}, err
	}

	builder.SetPort("8080")

	err = util.Add(
		builder,
		util.BuildAdd{From: mainExecutable, To: "/"},
	)
	if err != nil {
		return image.FunctionImage{}, err
	}

	builder.SetUser("1000")
	builder.SetCmd([]string{"/rust-faas"})
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
