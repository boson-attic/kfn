package js

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strings"

	"github.com/containers/image/types"
	"github.com/pkg/errors"

	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/image"
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"github.com/slinkydeveloper/kfn/pkg/util"
)

const (
	baseImage             = "oscf/js-runtime:0.0.2"
	npmScriptVariable     = "npm-script"
	devDependencyVariable = "dev-dependency"
	testFrameworkVariable = "test-framework"
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

	// Configure test file
	testFile := util.UnitTestFile(mainFile)
	if util.FsExist(testFile) {
		err = util.MkdirpIfNotExists(path.Join(targetDirectory, "usr", "test"))
		if err != nil {
			return err
		}

		err := util.Copy(testFile, path.Join(targetDirectory, "usr", "test", "index_test.js"))
		if err != nil {
			return err
		}
	}

	return util.Copy(runtimeDirectory(), path.Join(targetDirectory, "src"))
}

func (r jsLanguageManager) UnitTest(mainFile string, functionConfiguration map[string][]string, targetDirectory string) error {
	err := util.CommandsExists("npm")
	if err != nil {
		return err
	}

	// npm install --only=dev to install test framework
	err = util.RunCommand("npm", []string{"install", "--only=dev"}, path.Join(targetDirectory, "usr"))
	if err != nil {
		return err
	}

	err = util.RunCommand("npm", []string{"test"}, path.Join(targetDirectory, "usr"))
	if err != nil {
		return errors.Wrap(err, "error occurred while testing function")
	}

	return nil
}

func runtimeDirectory() string {
	return path.Join(config.RuntimeDir, "js")
}

func generatePackageJson(configurationEntries map[string][]string) ([]byte, error) {
	root := make(map[string]interface{})

	root["name"] = "function"
	root["version"] = "0.0.1"
	root["description"] = ""

	// First, apply the test-framework variable
	var tf string
	if tfArray, ok := configurationEntries[testFrameworkVariable]; ok && len(tfArray) == 1 {
		tf = strings.Trim(tfArray[0], " ")
	}
	for _, f := range getTestFrameworkPackageJsonConfig(tf) {
		f(root)
	}

	// Now let's apply custom deps
	depsRoot := util.GetNestedMap(root, "dependencies")
	devDepsRoot := util.GetNestedMap(root, "devDependencies")

	if deps, ok := configurationEntries[util.DEPENDENCY]; ok {
		for _, dep := range deps {
			splitted := strings.Split(dep, " ")
			if len(splitted) == 2 {
				depsRoot[strings.Trim(splitted[0], " ")] = strings.Trim(splitted[1], " ")
			} else if len(splitted) == 1 {
				pkgName := strings.Trim(splitted[0], " ")
				depsRoot[pkgName] = mustGetLatestNpmPackageVersionOrLatest(pkgName)
			} else {
				return nil, fmt.Errorf("invalid dependency entry: %v", dep)
			}
		}
	}

	if deps, ok := configurationEntries[devDependencyVariable]; ok {
		for _, dep := range deps {
			splitted := strings.Split(dep, " ")
			if len(splitted) == 2 {
				devDepsRoot[strings.Trim(splitted[0], " ")] = strings.Trim(splitted[1], " ")
			} else if len(splitted) == 1 {
				pkgName := strings.Trim(splitted[0], " ")
				devDepsRoot[pkgName] = mustGetLatestNpmPackageVersionOrLatest(pkgName)
			} else {
				return nil, fmt.Errorf("invalid dev dependency entry: %v", dep)
			}
		}
	}

	// Now let's check custom scripts
	scriptsRoot := util.GetNestedMap(root, "scripts")
	if scripts, ok := configurationEntries[npmScriptVariable]; ok {
		for _, script := range scripts {
			splitted := strings.SplitN(script, " ", 2)
			if len(splitted) == 2 {
				scriptsRoot[strings.Trim(splitted[0], " ")] = strings.Trim(splitted[1], " ")
			} else {
				return nil, fmt.Errorf("invalid script entry: %v", script)
			}
		}
	}
	return json.MarshalIndent(root, "", "  ")
}
