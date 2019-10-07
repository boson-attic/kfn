package pkg

import (
	"fmt"
	"github.com/containers/image/types"
	log "github.com/sirupsen/logrus"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/image"
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"github.com/slinkydeveloper/kfn/pkg/util"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Build function handles the different build steps:
// 1. Create target dir and runtime dir if not existing
// 2. Resolve the runtime manager
// 3. Download required runtime files if needed
// 4. Resolve the compiler manager
// 5. Check if compile dependencies are available (compiler, libraries, etc)
// 6. Download on the local filesystem the function if remote
// 7. Run compilation and output the files to move on the image
// 8. Resolve image builder
// 9. Build image and return the image id
func Build(location string, language languages.Language, imageName string, imageTag string, systemContext *types.SystemContext) (image.FunctionImage, error) {
	err := util.MkdirpIfNotExists(config.TargetDir)
	if err != nil {
		return image.FunctionImage{}, err
	}

	err = util.MkdirpIfNotExists(config.RuntimeDir)
	if err != nil {
		return image.FunctionImage{}, err
	}

	languageManager := languages.ResolveLanguageManager(language)

	err = languageManager.DownloadRuntimeIfRequired()
	if err != nil {
		return image.FunctionImage{}, err
	}

	log.Info("Checking compile dependencies")

	err = languageManager.CheckCompileDependencies()
	if err != nil {
		return image.FunctionImage{}, err
	}

	if strings.HasPrefix(location, "http") {
		log.Infof("Downloading function from %s", location)

		location, err = downloadFunctionFromHTTP(location, languages.GetExtension(language))
		if err != nil {
			return image.FunctionImage{}, err
		}
	} else {
		if !util.FileExist(location) {
			return image.FunctionImage{}, fmt.Errorf("cannot find file %s", location)
		}
	}

	log.Info("Configuring target directory")

	err = languageManager.ConfigureTargetDirectory(location, false)
	if err != nil {
		return image.FunctionImage{}, err
	}

	log.Info("Compiling")

	compiledOutput, additionalFiles, err := languageManager.Compile(location)
	if err != nil {
		return image.FunctionImage{}, err
	}

	log.Info("Starting build image")

	return languageManager.BuildImage(systemContext, imageName, imageTag, compiledOutput, additionalFiles)
}

func downloadFunctionFromHTTP(remote, extension string) (string, error) {
	f, err := ioutil.TempFile("", fmt.Sprintf("*.%d", extension))
	if err != nil {
		return "", err
	}

	err = util.DownloadAndPipe(remote, f)
	if err != nil {
		return "", err
	}
	err = f.Close()
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}
