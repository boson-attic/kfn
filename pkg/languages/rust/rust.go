package rust

import (
	"context"
	"fmt"
	"github.com/containers/image/types"
	"github.com/sirupsen/logrus"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/image"
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"github.com/slinkydeveloper/kfn/pkg/util"
	"io/ioutil"
	"os/exec"
	"path"
)

const (
	rustRuntimeRemoteZip = "https://github.com/openshift-cloud-functions/faas-rust-runtime/archive/master.zip"
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

	packageJsonTempl, err := r.resourceLoader.LoadTemplate("Cargo.toml")
	if err != nil {
		return err
	}

	err = util.PipeTemplateToFile(path.Join(targetDirectory, "Cargo.toml"), packageJsonTempl, struct {
		FunctionName string
	}{functionName})
	if err != nil {
		return err
	}

	return util.WriteFiles(
		targetDirectory,
		util.WriteDest{fmt.Sprintf("%s.rs", functionName), main},
	)
}

func (r rustLanguageManager) DownloadRuntimeIfRequired() error {
	if !util.FsExist(RuntimeDirectory()) {
		if err := util.MkdirpIfNotExists(RuntimeDirectory()); err != nil {
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

		if err := util.Copy(path.Join(tempDir, "faas-rust-runtime-master", "src"), path.Join(RuntimeDirectory(), "src")); err != nil {
			return err
		}

		if err := util.Copy(path.Join(tempDir, "faas-rust-runtime-master", "Cargo.toml"), path.Join(RuntimeDirectory())); err != nil {
			return err
		}

		logrus.Infof("Copied runtime to %s", RuntimeDirectory())
	} else {
		logrus.Infof("Using runtime cached in %s", RuntimeDirectory())
	}
	return nil
}

func (r rustLanguageManager) CheckCompileDependencies() error {
	return util.CommandsExists("rustc", "cargo")
}

func (r rustLanguageManager) ConfigureEditingDirectory(mainFile string) (string, string, error) {
	initialPath := path.Dir(mainFile)

	cargoDescriptor := path.Join(config.EditingDir, "Cargo.toml")
	functionFile := path.Join(config.EditingDir, "lib.rs")

	err := util.Link(path.Join(initialPath, "Cargo.toml"), cargoDescriptor)
	if err != nil {
		return "", "", err
	}

	err = util.Link(mainFile, functionFile)
	if err != nil {
		return "", "", err
	}

	return config.EditingDir, cargoDescriptor, nil
}

func (r rustLanguageManager) ConfigureTargetDirectory(mainFile string) error {
	if !util.FsExist(path.Join(path.Dir(mainFile), "Cargo.toml")) {
		return fmt.Errorf("Cannot find Cargo.toml in %s", path.Dir(mainFile))
	}

	if err := util.MkdirpIfNotExists(path.Join(config.TargetDir, "function")); err != nil {
		return err
	}

	err := util.Copy(mainFile, path.Join(config.TargetDir, "function", "lib.rs"))
	if err != nil {
		return err
	}

	err = util.Copy(path.Join(path.Dir(mainFile), "Cargo.toml"), path.Join(config.TargetDir, "function", "Cargo.toml"))
	if err != nil {
		return err
	}

	return util.CopyContent(RuntimeDirectory(), path.Join(config.TargetDir, "runtime"))
}

func (r rustLanguageManager) Compile(inputFile string) (string, []string, error) {
	compileCommand := exec.Command("cargo", "build") // "--release"
	// Root Cargo.toml is in runtime dir in runtime
	compileCommand.Dir = path.Join(config.TargetDir, "runtime")
	// Configure proper logging
	compileCommand.Stdout = config.GetLoggerWriter()
	compileCommand.Stderr = config.GetLoggerWriter()

	err := compileCommand.Run()
	if err != nil {
		return "", nil, err
	}

	return path.Join(config.TargetDir, "runtime", "target", "debug" /* release */, "rust-faas"), nil, nil
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
	return rustLanguageManager{util.NewResourceLoader("../../bootstrap_templates/rust")}
}

func RuntimeDirectory() string {
	return path.Join(config.RuntimeDir, "rust")
}


