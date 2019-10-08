package config

import (
	"fmt"
	"github.com/containers/buildah"
	"github.com/containers/buildah/pkg/parse"
	"github.com/containers/image/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
)

const (
	TARGET_DIR_ENV      = "target_dir"
	RUNTIME_DIR_ENV     = "runtime_dir"
	EDITING_DIR_ENV		= "editing_dir"
	VERBOSE             = "verbose"
	REGISTRY            = "registry"
	REGISTRY_USERNAME   = "registry_username"
	REGISTRY_PASSWORD   = "registry_password"
	REGISTRY_TLS_VERIFY = "registry_tls_verify"
	KUBECONFIG          = "kubeconfig"
	DEBUG               = "kfn_debug"
	CONFIG				= "config"
	NAMESPACE			= "namespace"
)

const (
	targetDirBase  string = "target"
	runtimeDirBase string = "runtime"
	editingDirBase string = "editing"
)

var (
	TargetDir              string
	RuntimeDir             string
	EditingDir			   string
	Verbose                bool
	Debug                  bool
	ImageRegistry          string
	ImageRegistryUsername  string
	ImageRegistryPassword  string
	ImageRegistryTLSVerify bool
	Kubeconfig             string
	BuildahIsolation       buildah.Isolation
	BuildSystemContext     *types.SystemContext
)

func InitLogging() {
	Verbose = getEnvBoolOrDefault(VERBOSE, false)
	Debug = getEnvBoolOrDefault(DEBUG, false)
	if Debug {
		log.SetLevel(log.DebugLevel)
	} else if Verbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

func InitDirVariables() {
	wd, _ := os.Getwd()

	// Configure variables
	TargetDir = path.Join(wd, getEnvStringOrDefault(TARGET_DIR_ENV, targetDirBase))
	RuntimeDir = path.Join(wd, getEnvStringOrDefault(RUNTIME_DIR_ENV, runtimeDirBase))
	EditingDir = path.Join(wd, getEnvStringOrDefault(EDITING_DIR_ENV, editingDirBase))
}

func InitBuildVariables(cmd *cobra.Command) error {
	ImageRegistry = getEnvOrFail(REGISTRY)
	ImageRegistryUsername = getEnvStringOrDefault(REGISTRY_USERNAME, "")
	ImageRegistryPassword = getEnvStringOrDefault(REGISTRY_PASSWORD, "")
	ImageRegistryTLSVerify = getEnvBoolOrDefault(REGISTRY_TLS_VERIFY, true)

	var err error
	BuildSystemContext, err = parseSystemContext(cmd)
	ImageRegistry, ImageRegistryUsername, ImageRegistryPassword, err = inferImageRegistry(BuildSystemContext)
	if err != nil {
		return err
	}
	setSystemContextCredentials(BuildSystemContext, ImageRegistryUsername, ImageRegistryPassword)

	BuildahIsolation = getBuildahIsolation()

	return nil
}

func InitRunVariables() {
	Kubeconfig = getEnvStringOrDefault(KUBECONFIG, "")
}

func getBuildahIsolation() buildah.Isolation {
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

func parseSystemContext(cmd *cobra.Command) (*types.SystemContext, error) {
	ctx, err := parse.SystemContextFromOptions(cmd)
	if err != nil {
		return nil, err
	}

	if ImageRegistryUsername != "" {
		ctx.DockerAuthConfig = &types.DockerAuthConfig{
			Username: ImageRegistryUsername,
			Password: ImageRegistryPassword,
		}
	}

	ctx.DockerInsecureSkipTLSVerify = types.NewOptionalBool(!ImageRegistryTLSVerify)
	ctx.OCIInsecureSkipTLSVerify = !ImageRegistryTLSVerify
	ctx.DockerDaemonInsecureSkipTLSVerify = !ImageRegistryTLSVerify

	return ctx, nil
}

func setSystemContextCredentials(sysContext *types.SystemContext, username, password string) {
	if username != "" {
		sysContext.DockerAuthConfig = &types.DockerAuthConfig{
			Username: username,
			Password: password,
		}
	}
}
