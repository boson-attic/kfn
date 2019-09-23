package pkg

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
	"path"
)

const (
	TARGET_DIR_ENV    = "target_dir"
	RUNTIME_DIR_ENV   = "runtime_dir"
	VERBOSE           = "verbose"
	REGISTRY          = "registry"
	REGISTRY_USERNAME = "registry_username"
	REGISTRY_PASSWORD = "registry_password"
	KUBECONFIG        = "kubeconfig"
)

const (
	targetDirBase  string = "target"
	runtimeDirBase string = "runtime"
)

var (
	TargetDir             string
	RuntimeDir            string
	Verbose               bool
	ImageRegistry         string
	ImageRegistryUsername string
	ImageRegistryPassword string
	Kubeconfig            string
)

func InitVariables() {
	wd, _ := os.Getwd()

	// Configure variables
	TargetDir = path.Join(wd, getEnvOrDefault(TARGET_DIR_ENV, targetDirBase))
	RuntimeDir = path.Join(wd, getEnvOrDefault(RUNTIME_DIR_ENV, runtimeDirBase))
	Verbose = getEnvBoolOrDefault(VERBOSE, false)
	ImageRegistry = getEnvOrFail(REGISTRY)
	ImageRegistryUsername = getEnvOrDefault(REGISTRY_USERNAME, "")
	ImageRegistryPassword = getEnvOrDefault(REGISTRY_PASSWORD, "")
	Kubeconfig = getEnvOrDefault(KUBECONFIG, "")

	// Configure logging
	if Verbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

func GetLoggerWriter() *io.PipeWriter {
	logger := log.New()
	if Verbose {
		logger.SetLevel(log.InfoLevel)
	} else {
		logger.SetLevel(log.WarnLevel)
	}
	return logger.Writer()
}

func getEnvOrFail(envName string) string {
	if viper.IsSet(envName) {
		return viper.GetString(envName)
	} else {
		panic(fmt.Errorf("Cannot find flag/env/config entry %s", envName))
	}
}

func getEnvOrDefault(envName string, defaultValue string) string {
	if viper.IsSet(envName) {
		return viper.GetString(envName)
	} else {
		return defaultValue
	}
}

func getEnvBoolOrDefault(envName string, defaultValue bool) bool {
	if viper.IsSet(envName) {
		return viper.GetBool(envName)
	} else {
		return defaultValue
	}
}
