package js

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/containers/image/types"
	"github.com/sirupsen/logrus"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/image"
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"github.com/slinkydeveloper/kfn/pkg/util"
)

const (
	baseImage          = "node:12-alpine"
	jsRuntimeRemoteZip = "https://github.com/openshift-cloud-functions/faas-js-runtime-image/archive/master.zip"
)

type jsLanguageManager struct {
	resourceLoader util.ResourceLoader
}

func NewJsLanguageManger() languages.LanguageManager {
	return jsLanguageManager{util.NewResourceLoader("../../templates/js")}
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

	return util.WriteFiles(
		targetDirectory,
		util.WriteDest{fmt.Sprintf("%s.js", functionName), main},
	)
}

// For JS Runtime nothing is required, since the building atm is done directly in the container
func (j jsLanguageManager) CheckCompileDependencies() error {
	return nil
}

func (j jsLanguageManager) Compile(mainFile string, functionConfiguration map[string][]string, targetDirectory string) (string, []string, error) {
	dir, _ := path.Split(mainFile)
	packageJson := path.Join(dir, "package.json")
	if util.FsExist(packageJson) {
		return mainFile, []string{packageJson}, nil
	} else {
		return mainFile, []string{}, nil
	}
}

func (j jsLanguageManager) BuildImage(systemContext *types.SystemContext, imageName string, imageTag string, mainExecutable string, additionalFiles []string, targetDirectory string) (image.FunctionImage, error) {
	builder, err := util.InitializeBuilder(context.TODO(), systemContext, baseImage)
	if err != nil {
		return image.FunctionImage{}, err
	}

	builder.SetPort("8080")

	err = util.Add(builder, util.BuildAdd{From: path.Join(targetDirectory, "src"), To: "/home/node/src"}, util.BuildAdd{From: path.Join(targetDirectory, "usr"), To: "/home/node/usr"})
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
	if !util.FsExist(runtimeDirectory()) {
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

		if err := util.CopyContent(path.Join(tempDir, "faas-js-runtime-image-master", "src"), runtimeDirectory()); err != nil {
			return err
		}

		logrus.Infof("Copied runtime to %s", runtimeDirectory())
	} else {
		logrus.Infof("Using runtime cached in %s", runtimeDirectory())
	}
	return nil
}

func (r jsLanguageManager) ConfigureEditingDirectory(mainFile string, functionConfiguration map[string][]string, editingDirectory string) (string, error) {
	functionFile := path.Join(editingDirectory, "index.js")
	err := util.Link(mainFile, functionFile)
	if err != nil {
		return "", err
	}

	packageJson, err := generatePackageJson(functionConfiguration)
	if err != nil {
		return "", err
	}

	err = util.WriteFiles(editingDirectory, util.WriteDest{Filename: "package.json", Data: packageJson})
	if err != nil {
		return "", err
	}

	return editingDirectory, nil
}

func (j jsLanguageManager) ConfigureTargetDirectory(mainFile string, functionConfiguration map[string][]string, targetDirectory string) error {
	if err := util.MkdirpIfNotExists(path.Join(targetDirectory, "usr")); err != nil {
		return err
	}

	err := util.Copy(mainFile, path.Join(targetDirectory, "usr", "index.js"))
	if err != nil {
		return err
	}

	packageJson, err := generatePackageJson(functionConfiguration)
	if err != nil {
		return err
	}

	err = util.WriteFiles(path.Join(targetDirectory, "usr"), util.WriteDest{Filename: "package.json", Data: packageJson})
	if err != nil {
		return err
	}

	return util.Copy(runtimeDirectory(), path.Join(targetDirectory, "src"))
}

func runtimeDirectory() string {
	return path.Join(config.RuntimeDir, "js")
}

func generatePackageJson(configurationEntries map[string][]string) ([]byte, error) {
	root := make(map[string]interface{})

	root["name"] = "function"
	root["version"] = "0.0.1"
	root["description"] = ""

	depsRoot := make(map[string]string)

	if deps, ok := configurationEntries[util.DEPENDENCY]; ok {
		for _, dep := range deps {
			splitted := strings.Split(dep, " ")
			if len(splitted) != 2 {
				return nil, fmt.Errorf("Invalid dependency entry: %v", dep)
			}
			depsRoot[strings.Trim(splitted[0], " ")] = strings.Trim(splitted[1], " ")
		}
	}

	root["dependencies"] = depsRoot

	return json.MarshalIndent(root, "", "  ")
}
