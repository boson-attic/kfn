package js

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg/kfn/util"
	"path"
)

type JsBoostrapper struct {
	resourceLoader util.ResourceLoader
}

func NewJsBootstrapper() JsBoostrapper {
	return JsBoostrapper{resourceLoader: util.NewResourceLoader("../../../bootstrap_templates/js")}
}

func (r JsBoostrapper) Bootstrap(functionName string, targetDirectory string) error {
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
