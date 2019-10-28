package languages

import (
	"github.com/containers/image/types"
	"github.com/slinkydeveloper/kfn/pkg/image"
)

type Language uint8

func (l Language) Bootstrap(functionName string, targetDirectory string) error {
	return ResolveLanguageManager(l).Bootstrap(functionName, targetDirectory)
}

func (l Language) CheckCompileDependencies() error {
	return ResolveLanguageManager(l).CheckCompileDependencies()
}

func (l Language) UnitTest(mainFile string, functionConfiguration map[string][]string, targetDirectory string) error {
	return ResolveLanguageManager(l).UnitTest(mainFile, functionConfiguration, targetDirectory)
}

func (l Language) Compile(mainFile string, functionConfiguration map[string][]string, targetDirectory string) (string, []string, error) {
	return ResolveLanguageManager(l).Compile(mainFile, functionConfiguration, targetDirectory)
}

func (l Language) DownloadRuntimeIfRequired() error {
	return ResolveLanguageManager(l).DownloadRuntimeIfRequired()
}

func (l Language) ConfigureEditingDirectory(mainFile string, functionConfiguration map[string][]string, editingDirectory string) (string, error) {
	return ResolveLanguageManager(l).ConfigureEditingDirectory(mainFile, functionConfiguration, editingDirectory)
}

func (l Language) ConfigureTargetDirectory(mainFile string, functionConfiguration map[string][]string, targetDirectory string) error {
	return ResolveLanguageManager(l).ConfigureTargetDirectory(mainFile, functionConfiguration, targetDirectory)
}

func (l Language) BuildImage(systemContext *types.SystemContext, imageName string, imageTag string, mainExecutable string, additionalFiles []string, targetDirectory string) (image.FunctionImage, error) {
	return ResolveLanguageManager(l).BuildImage(systemContext, imageName, imageTag, mainExecutable, additionalFiles, targetDirectory)
}

const (
	Javascript Language = iota
	Rust
	Unknown
)

func GetExtension(language Language) string {
	switch language {
	case Javascript:
		return "js"
	case Rust:
		return "rs"
	default:
		return ""
	}
}

func GetLineComment(language Language) string {
	return "//"
}

func GetLanguage(ext string) Language {
	switch ext {
	case ".js":
		return Javascript
	case ".rs":
		return Rust
	default:
		return Unknown
	}
}
