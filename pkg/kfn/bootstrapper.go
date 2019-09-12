package kfn

import (
	"github.com/slinkydeveloper/kfn/pkg/kfn/js"
)

type Bootstrapper interface {
	Bootstrap(functionName string, targetDirectory string) error
}

var (
	jsBootstrapper Bootstrapper = js.NewJsBootstrapper()
)

func ResolveBootrapper(language Language) *Bootstrapper {
	switch language {
	case Javascript:
		return &jsBootstrapper
	default:
		return nil
	}
}
