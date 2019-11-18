package pkg

import (
	"time"

	"github.com/containers/image/types"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/rand"

	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/image"
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"github.com/slinkydeveloper/kfn/pkg/util"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func PrepareBuildEnvironment(location string, language languages.Language) (targetDir string, functionConfiguration map[string][]string, err error) {
	targetDir = config.GetTargetDir(location)

	err = util.MkdirpIfNotExists(targetDir)
	if err != nil {
		return
	}

	err = util.MkdirpIfNotExists(config.RuntimeDir)
	if err != nil {
		return
	}

	err = language.DownloadRuntimeIfRequired()
	if err != nil {
		return
	}

	err = language.CheckCompileDependencies()
	if err != nil {
		return
	}

	log.Infof("Retrieving function configuration")

	functionConfiguration, err = util.ParseConfigComments(languages.GetLineComment(language), location)
	if err != nil {
		return
	}

	// Log only if needed
	if config.Verbose {
		for k, v := range functionConfiguration {
			log.Infof("Configuration entry %s: %s", k, v)
		}
	}

	log.Info("Configuring target directory")

	err = language.ConfigureTargetDirectory(location, functionConfiguration, targetDir)
	if err != nil {
		return
	}

	return
}

func Build(location string, language languages.Language, imageName string, imageTag string, systemContext *types.SystemContext) (image.FunctionImage, error) {
	targetDir, functionConfiguration, err := PrepareBuildEnvironment(location, language)
	if err != nil {
		return image.FunctionImage{}, err
	}

	log.Info("Compiling")

	compiledOutput, additionalFiles, err := language.Compile(location, functionConfiguration, targetDir)
	if err != nil {
		return image.FunctionImage{}, err
	}

	log.Info("Starting build image")

	return language.BuildImage(systemContext, imageName, imageTag, compiledOutput, additionalFiles, targetDir)
}

func Test(location string, language languages.Language) error {
	targetDir, functionConfiguration, err := PrepareBuildEnvironment(location, language)
	if err != nil {
		return err
	}

	return language.UnitTest(location, functionConfiguration, targetDir)
}
