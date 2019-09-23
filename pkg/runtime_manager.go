package pkg

import "github.com/slinkydeveloper/kfn/pkg/js"

type RuntimeManager interface {
	DownloadRuntimeIfRequired() error
	ConfigureTargetDirectory(mainFile string, additionalFiles []string) error
}

var (
	jsRuntimeManager = js.NewJSRuntimeManager()
)

func ResolveRuntimeManager(language Language) *RuntimeManager {
	switch language {
	case Javascript:
		return &jsRuntimeManager
	default:
		return nil
	}
}
