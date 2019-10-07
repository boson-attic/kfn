package js

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
	"path"
)

const (
	baseImage          = "node:12-alpine"
	jsRuntimeRemoteZip = "https://github.com/openshift-cloud-functions/faas-js-runtime-image/archive/master.zip"
)

type jsLanguageManager struct {
	resourceLoader util.ResourceLoader
}

func NewJsLanguageManger() languages.LanguageManager {
	return jsLanguageManager{util.NewResourceLoader("../../bootstrap_templates/js")}
}

func (r jsLanguageManager) Bootstrap(functionName string, targetDirectory string) error {
	err := util.MkdirpIfNotExists(targetDirectory)
	if err != nil {
		return err
	}

	main, err := r.resourceLoader.LoadResource("index.js")
	if err != nil {
		return err
	}

	packageJsonTempl, err := r.resourceLoader.LoadTemplate("package.json")
	if err != nil {
		return err
	}

	err = util.PipeTemplateToFile(path.Join(targetDirectory, "package.json"), packageJsonTempl, struct {
		FunctionName string
	}{functionName})
	if err != nil {
		return err
	}

	return util.WriteFiles(
		targetDirectory,
		util.WriteDest{fmt.Sprintf("%s.js", functionName), main},
	)
}

// For JS Runtime nothing is required, since the building atm is done directly in the container
func (j jsLanguageManager) CheckCompileDependencies() error {
	return nil
}

func (j jsLanguageManager) Compile(inputFile string) (string, []string, error) {
	dir, _ := path.Split(inputFile)
	packageJson := path.Join(dir, "package.json")
	if util.FsExist(packageJson) {
		return inputFile, []string{packageJson}, nil
	} else {
		return inputFile, []string{}, nil
	}
}

func (j jsLanguageManager) BuildImage(systemContext *types.SystemContext, imageName string, imageTag string, mainExecutable string, additionalFiles []string) (image.FunctionImage, error) {
	builder, err := util.InitializeBuilder(context.TODO(), systemContext, baseImage)
	if err != nil {
		return image.FunctionImage{}, err
	}

	builder.SetPort("8080")

	err = util.Add(builder, util.BuildAdd{From: path.Join(config.TargetDir, "src"), To: "/home/node/src"}, util.BuildAdd{From: path.Join(config.TargetDir, "usr"), To: "/home/node/usr"})
	if err != nil {
		return image.FunctionImage{}, err
	}

	err = util.RunCommands(
		builder,
		util.BuildCommand{Command: "mkdir -p /home/node/usr/.npm"},
		util.BuildCommand{Command: "chmod -R a+g+x /home/node/usr"},
		util.BuildCommand{Command: "chmod -R a+g+x /home/node/src"},
		util.BuildCommand{"npm install", "/home/node/usr"},
		util.BuildCommand{"npm install", "/home/node/src"},
	)
	if err != nil {
		return image.FunctionImage{}, err
	}

	builder.SetEnv("HOME", "/home/node/usr")
	builder.SetUser("1000")
	builder.SetWorkDir("/home/node/src")

	builder.SetCmd([]string{"node", "/home/node/src/index.js"})

	return util.CommitImage(builder, systemContext, imageName, imageTag)
}

func (j jsLanguageManager) DownloadRuntimeIfRequired() error {
	if !util.FsExist(RuntimeDirectory()) {
		if err := util.MkdirpIfNotExists(config.RuntimeDir); err != nil {
			return err
		}

		tempDir, err := ioutil.TempDir("", "faas-js-runtime-image")
		if err != nil {
			return err
		}

		runtimeZip := path.Join(tempDir, "master.zip")

		if err := util.DownloadFile(jsRuntimeRemoteZip, runtimeZip); err != nil {
			return err
		}

		logrus.Infof("Downloading runtime from %s to %s", jsRuntimeRemoteZip, runtimeZip)

		if _, err := util.Unzip(runtimeZip, tempDir); err != nil {
			return err
		}

		logrus.Infof("Runtime unzipped to %s", tempDir)

		if err := util.CopyContent(path.Join(tempDir, "faas-js-runtime-image-master", "src"), RuntimeDirectory()); err != nil {
			return err
		}

		logrus.Infof("Copied runtime to %s", RuntimeDirectory())
	} else {
		logrus.Infof("Using runtime cached in %s", RuntimeDirectory())
	}
	return nil
}

func (j jsLanguageManager) ConfigureTargetDirectory(mainFile string, linkOnly bool) error {
	if err := util.MkdirpIfNotExists(path.Join(config.TargetDir, "usr")); err != nil {
		return err
	}

	cp := util.CopyOrLink(linkOnly)

	err := cp(mainFile, path.Join(config.TargetDir, "usr", "index.js"))
	if err != nil {
		return err
	}

	packageJsonPath := path.Join(path.Dir(mainFile), "package.json")
	if util.FsExist(packageJsonPath) {
		err = cp(packageJsonPath, path.Join(config.TargetDir, "usr"))
		if err != nil {
			return err
		}
	}

	return cp(path.Join(RuntimeDirectory()), path.Join(config.TargetDir, "src"))
}

func RuntimeDirectory() string {
	return path.Join(config.RuntimeDir, "js")
}
