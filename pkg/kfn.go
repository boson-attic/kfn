package pkg

import (
	"fmt"
	"github.com/containers/image/types"
	log "github.com/sirupsen/logrus"
	"github.com/slinkydeveloper/kfn/pkg/image"
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
func Build(location string, language Language, imageName string, imageTag string, systemContext *types.SystemContext) (image.FunctionImage, error) {
	err := util.MkdirpIfNotExists(TargetDir)
	if err != nil {
		return "", err
	}

	err = util.MkdirpIfNotExists(RuntimeDir)
	if err != nil {
		return "", err
	}

	runtimeManager := ResolveRuntimeManager(language)

	err = (*runtimeManager).DownloadRuntimeIfRequired()
	if err != nil {
		return "", err
	}

	compilerManager := ResolveCompilerManager(language)

	log.Info("Checking compile dependencies")

	err = (*compilerManager).CheckCompileDependencies()
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(location, "http") {
		log.Infof("Downloading function from %s", location)

		location, err = downloadFunctionFromHTTP(location, getExtension(language))
		if err != nil {
			return "", err
		}
	}

	log.Info("Compiling")

	main, additionalFiles, err := (*compilerManager).Compile(location)
	if err != nil {
		return "", err
	}

	log.Info("Configuring target directory")

	err = (*runtimeManager).ConfigureTargetDirectory(main, additionalFiles)
	if err != nil {
		return "", err
	}

	log.Info("Starting build image")

	imageBuilder := ResolveImageBuilder(language)
	return (*imageBuilder).BuildImage(systemContext, image.FunctionImage{ImageName: imageName, Tag: imageTag})
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
