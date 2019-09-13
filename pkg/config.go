package pkg

import (
	"github.com/spf13/viper"
	"os"
	"path"
)

const (
	TARGET_DIR_ENV  = "target_dir"
	RUNTIME_DIR_ENV = "runtime_dir"
)

const (
	targetDirBase  string = "target"
	runtimeDirBase string = "runtime"
)

var (
	TargetDir  string
	RuntimeDir string
)

func InitVariables() {
	wd, _ := os.Getwd()
	TargetDir = path.Join(wd, getEnvOrDefault(TARGET_DIR_ENV, targetDirBase))
	RuntimeDir = path.Join(wd, getEnvOrDefault(RUNTIME_DIR_ENV, runtimeDirBase))
}

func getEnvOrDefault(envName string, defaultValue string) string {
	if viper.IsSet(envName) {
		return viper.GetString(envName)
	} else {
		return defaultValue
	}
}
