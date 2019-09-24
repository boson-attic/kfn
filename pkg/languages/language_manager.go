package languages

import (
	"github.com/containers/image/types"
	"github.com/slinkydeveloper/kfn/pkg/image"
)

type LanguageManager interface {
	// Bootstrap a new function
	Bootstrap(functionName string, targetDirectory string) error

	// Check if compiler/package manager are installed
	CheckCompileDependencies() error

	// Compile with Main input file, additional files
	Compile(inputFile string) (string, []string, error)

	// Download the runtime required to build the function
	DownloadRuntimeIfRequired() error

	// Configure target directory with symbolic links
	ConfigureTargetDirectory(mainFile string, additionalFiles []string, linkOnly bool) error

	// Build the container image
	BuildImage(systemContext *types.SystemContext, imageName string, imageTag string) (image.FunctionImage, error)
}

var (
	languageManagerMap = make(map[Language]LanguageManager, 0)
)

func RegisterLanguageManager(language Language, manager LanguageManager) {
	languageManagerMap[language] = manager
}

func ResolveLanguageManager(language Language) LanguageManager {
	return languageManagerMap[language]
}
