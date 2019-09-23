package js

import "github.com/slinkydeveloper/kfn/pkg"

type jsCompilerManager struct{}

func NewJSCompilerManager() pkg.CompilerManager {
	return jsCompilerManager{}
}

// For JS Runtime nothing is required, since the building atm is done directly in the container
func (j jsCompilerManager) CheckCompileDependencies() error {
	return nil
}

func (j jsCompilerManager) Compile(inputFile string) (string, []string, error) {
	return inputFile, []string{}, nil
}
