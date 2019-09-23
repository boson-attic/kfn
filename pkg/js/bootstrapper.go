package js

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg"
	"github.com/slinkydeveloper/kfn/pkg/util"
	"path"
)

type jsBoostrapper struct {
	resourceLoader util.ResourceLoader
}

func NewJSBootstrapper() pkg.Bootstrapper {
	return jsBoostrapper{resourceLoader: util.NewResourceLoader("../../../bootstrap_templates/js")}
}

func (r jsBoostrapper) Bootstrap(functionName string, targetDirectory string) error {
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
