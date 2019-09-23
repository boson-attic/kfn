package pkg

import "github.com/slinkydeveloper/kfn/pkg/js"

type CompilerManager interface {
	// Check if compiler/package manager are installed
	CheckCompileDependencies() error

	// Main file, additional files
	Compile(inputFile string) (string, []string, error)
}

var (
	jsCompilerManager = js.NewJSCompilerManager()
)

func ResolveCompilerManager(language Language) *CompilerManager {
	switch language {
	case Javascript:
		return &jsCompilerManager
	default:
		return nil
	}
}
