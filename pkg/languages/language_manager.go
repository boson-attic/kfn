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

	// Download the runtime required to build the function
	DownloadRuntimeIfRequired() error

	// Configure a temp directory with symlinks required to edit the file
	ConfigureEditingDirectory(mainFile string, functionConfiguration map[string][]string, editingDirectory string) (directory string, err error)

	// Configure target directory with main file and eventual test file
	ConfigureTargetDirectory(mainFile string, functionConfiguration map[string][]string, targetDirectory string) error

	// Run unit tests
	UnitTest(mainFile string, functionConfiguration map[string][]string, targetDirectory string) error

	// Compile with Main input file, returns executable + additional files to copy
	Compile(mainFile string, functionConfiguration map[string][]string, targetDirectory string) (mainExecutable string, additionalFiles []string, err error)

	// Build the container image
	BuildImage(systemContext *types.SystemContext, imageName string, imageTag string, mainExecutable string, additionalFiles []string, targetDirectory string) (image.FunctionImage, error)
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
