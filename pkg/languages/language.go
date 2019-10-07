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

func (l Language) Compile(inputFile string) (string, []string, error) {
	return ResolveLanguageManager(l).Compile(inputFile)
}

func (l Language) DownloadRuntimeIfRequired() error {
	return ResolveLanguageManager(l).DownloadRuntimeIfRequired()
}

func (l Language) ConfigureTargetDirectory(mainFile string, linkOnly bool) error {
	return ResolveLanguageManager(l).ConfigureTargetDirectory(mainFile, linkOnly)
}

func (l Language) BuildImage(systemContext *types.SystemContext, imageName string, imageTag string, mainExecutable string, additionalFiles []string) (image.FunctionImage, error) {
	return ResolveLanguageManager(l).BuildImage(systemContext, imageName, imageTag, mainExecutable, additionalFiles)
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
