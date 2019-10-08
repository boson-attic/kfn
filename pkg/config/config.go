package config

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/containers/buildah"
	"github.com/containers/buildah/pkg/parse"
	"github.com/containers/image/types"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
)

const (
	KFN_DIR_ENV         = "kfn_dir"
	VERBOSE             = "verbose"
	REGISTRY            = "registry"
	REGISTRY_USERNAME   = "registry_username"
	REGISTRY_PASSWORD   = "registry_password"
	REGISTRY_TLS_VERIFY = "registry_tls_verify"
	KUBECONFIG          = "kubeconfig"
	DEBUG               = "kfn_debug"
	CONFIG              = "config"
	NAMESPACE           = "namespace"
)

const (
	kfnDirBase     string = ".kfn"
	targetDirBase  string = "target"
	runtimeDirBase string = "runtime"
	editingDirBase string = "editing"
)

var (
	KfnDir                 string
	Verbose                bool
	RuntimeDir             string
	Debug                  bool
	ImageRegistry          string
	ImageRegistryUsername  string
	ImageRegistryPassword  string
	ImageRegistryTLSVerify bool
	Kubeconfig             string
	Namespace              string
	BuildahIsolation       buildah.Isolation
	BuildSystemContext     *types.SystemContext
)

func init() {

}

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

func GetKfnDir() string {
	KfnDir = getEnvStringOrDefault(KFN_DIR_ENV, kfnDirBase)
	home, _ := homedir.Dir()
	if !filepath.IsAbs(KfnDir) {
		KfnDir = path.Join(home, KfnDir)
	}

	return KfnDir
}

func InitKfnDir() {
	KfnDir = getEnvStringOrDefault(KFN_DIR_ENV, kfnDirBase)

	// If filepath is absolute, leave as is
	if !filepath.IsAbs(KfnDir) {
		home, err := homedir.Dir()

		// Try to resolve starting from home dir
		if err == nil {
			KfnDir = path.Join(home, KfnDir)
		} else {
			// Cannot resolve home dir (it can happen in container environments), start from pwd
			KfnDir, err = filepath.Abs(KfnDir)
			if err != nil {
				panic(err)
			}
		}
	}

	RuntimeDir = path.Join(KfnDir, runtimeDirBase)

	log.Debugf("Kfn dir: %s", KfnDir)
}

func GetTargetDir(functionLocation string) string {
	return path.Join(KfnDir, getFunctionHash(functionLocation), targetDirBase)
}

func GetEditingDir(functionLocation string) string {
	return path.Join(KfnDir, getFunctionHash(functionLocation), editingDirBase)
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
	Namespace = getEnvStringOrDefault(NAMESPACE, "default")
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

func getFunctionHash(functionLocation string) string {
	return hex.EncodeToString(sha256.New().Sum([]byte(functionLocation)))
}
