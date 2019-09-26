package config

import (
	"fmt"
	"github.com/containers/buildah"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

const (
	TARGET_DIR_ENV      = "target_dir"
	RUNTIME_DIR_ENV     = "runtime_dir"
	VERBOSE             = "verbose"
	REGISTRY            = "registry"
	REGISTRY_USERNAME   = "registry_username"
	REGISTRY_PASSWORD   = "registry_password"
	REGISTRY_TLS_VERIFY = "registry_tls_verify"
	KUBECONFIG          = "kubeconfig"
	DEBUG               = "kfn_debug"
)

const (
	targetDirBase  string = "target"
	runtimeDirBase string = "runtime"
)

var (
	TargetDir              string
	RuntimeDir             string
	Verbose                bool
	ImageRegistry          string
	ImageRegistryUsername  string
	ImageRegistryPassword  string
	ImageRegistryTLSVerify bool
	Kubeconfig             string
)

func InitVariables() {
	wd, _ := os.Getwd()

	// Configure variables
	TargetDir = path.Join(wd, getEnvStringOrDefault(TARGET_DIR_ENV, targetDirBase))
	RuntimeDir = path.Join(wd, getEnvStringOrDefault(RUNTIME_DIR_ENV, runtimeDirBase))
	Verbose = getEnvBoolOrDefault(VERBOSE, false)
	ImageRegistry = getEnvOrFail(REGISTRY)
	ImageRegistryUsername = getEnvStringOrDefault(REGISTRY_USERNAME, "")
	ImageRegistryPassword = getEnvStringOrDefault(REGISTRY_PASSWORD, "")
	ImageRegistryTLSVerify = getEnvBoolOrDefault(REGISTRY_TLS_VERIFY, true)
	Kubeconfig = getEnvStringOrDefault(KUBECONFIG, "")

	// Configure logging
	if getEnvBoolOrDefault(DEBUG, false) {
		log.SetLevel(log.DebugLevel)
	} else if Verbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	var err error
	ImageRegistry, ImageRegistryUsername, ImageRegistryPassword, err = inferImageRegistry()
	if err != nil {
		panic(err)
	}

	log.Infof("Using Kubeconfig: %s", Kubeconfig)
	log.Infof("Using image registry: %s", ImageRegistry)
}

func GetBuildahIsolation() buildah.Isolation {
	var isolation buildah.Isolation

	envIsolation := os.Getenv("BUILDAH_ISOLATION")
	switch envIsolation {
	case "chroot":
		isolation = buildah.IsolationChroot
	default:
		isolation = buildah.IsolationOCIRootless
	}

	return isolation
}

func getEnvOrFail(envName string) string {
	if viper.IsSet(envName) {
		return viper.GetString(envName)
	} else {
		panic(fmt.Errorf("Cannot find flag/env/config entry %s", envName))
	}
}

func getEnvStringOrDefault(envName string, defaultValue string) string {
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
