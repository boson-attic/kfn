package js

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strings"

	"github.com/containers/image/types"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/image"
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"github.com/slinkydeveloper/kfn/pkg/util"
)

const (
	baseImage = "oscf/js-runtime:0.0.2"
)

type jsLanguageManager struct {
	resourceLoader util.ResourceLoader
}

// NewJsLanguageManger returns a new instance of a jsLanguageManager
func NewJsLanguageManger() languages.LanguageManager {
	return jsLanguageManager{util.NewResourceLoader("../../templates/js")}
}

func (j jsLanguageManager) Bootstrap(functionName string, targetDirectory string) error {
	err := util.MkdirpIfNotExists(targetDirectory)
	if err != nil {
		return err
	}

	main, err := j.resourceLoader.LoadResource("index.js")
	if err != nil {
		return err
	}

	return util.WriteFiles(
		targetDirectory,
		util.WriteDest{Filename: fmt.Sprintf("%s.js", functionName), Data: main},
	)
}

// CheckCompileDependencies always returns nil for Node.js
func (j jsLanguageManager) CheckCompileDependencies() error {
	return nil
}

func (j jsLanguageManager) Compile(mainFile string, functionConfiguration map[string][]string, targetDirectory string) (string, []string, error) {
	dir, _ := path.Split(mainFile)
	packageJSON := path.Join(dir, "package.json")
	if util.FsExist(packageJSON) {
		return mainFile, []string{packageJSON}, nil
	}
	return mainFile, []string{}, nil
}

func (j jsLanguageManager) BuildImage(systemContext *types.SystemContext, imageName string, imageTag string, mainExecutable string, additionalFiles []string, targetDirectory string) (image.FunctionImage, error) {
	builder, err := util.InitializeBuilder(context.TODO(), systemContext, baseImage)
	if err != nil {
		return image.FunctionImage{}, err
	}

	builder.SetPort("8080")

	err = util.Add(builder, util.BuildAdd{From: path.Join(targetDirectory, "usr"), To: "/home/node/usr"})
	if err != nil {
		return image.FunctionImage{}, err
	}

	err = util.RunCommands(
		builder,
		util.BuildCommand{Command: "npm install", Wd: "/home/node/usr"},
	)
	if err != nil {
		return image.FunctionImage{}, err
	}

	builder.SetEnv("HOME", "/home/node/usr")
	builder.SetUser("1001")
	builder.SetWorkDir("/home/node/src")

	builder.SetCmd([]string{"node", "/home/node/src/index.js"})

	return util.CommitImage(builder, systemContext, imageName, imageTag)
}

// DownloadRuntimeIfRequired is not used in the Node.js runtime
func (j jsLanguageManager) DownloadRuntimeIfRequired() error {
	return nil
}

func (r jsLanguageManager) ConfigureEditingDirectory(mainFile string, functionConfiguration map[string][]string, editingDirectory string) (string, error) {
	functionFile := path.Join(editingDirectory, "index.js")
	err := util.Link(mainFile, functionFile)
	if err != nil {
		return "", err
	}

	packageJSONFile := path.Join(path.Dir(mainFile), "package.json")
	packageJSONLink := path.Join(editingDirectory, "package.json")
	err = util.Link(packageJSONFile, packageJSONLink)
	if err != nil {
		return "", err
	}

	return editingDirectory, nil
}

func (j jsLanguageManager) ConfigureTargetDirectory(mainFile string, functionConfiguration map[string][]string, targetDirectory string) error {
	if err := util.MkdirpIfNotExists(path.Join(targetDirectory, "usr")); err != nil {
		return err
	}

	// Copy the user function file to the target directory
	err := util.Copy(mainFile, path.Join(targetDirectory, "usr", "index.js"))
	if err != nil {
		return err
	}

	// If a user package.json file exists, copy it.
	packageJSONFile := path.Join(path.Dir(mainFile), "package.json")
	if util.FileExist(packageJSONFile) {
		err = util.Copy(packageJSONFile, path.Join(targetDirectory, "usr", "package.json"))
		if err != nil {
			return err
		}
	} else {
		// No user package.json exists. Generate one.
		packageJSONData, err := generatePackageJSON(functionConfiguration)
		if err != nil {
			return err
		}
		err = util.WriteFiles(path.Join(targetDirectory, "usr"), util.WriteDest{Filename: "package.json", Data: packageJSONData})
		if err != nil {
			return err
		}
	}

	return util.Copy(runtimeDirectory(), path.Join(targetDirectory, "src"))
}

func runtimeDirectory() string {
	return path.Join(config.RuntimeDir, "js")
}

func generatePackageJSON(configurationEntries map[string][]string) ([]byte, error) {
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
