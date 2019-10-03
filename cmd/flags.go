package cmd

import (
	"github.com/containers/buildah/pkg/parse"
	"github.com/containers/image/types"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

var (
	imageName string
	imageTag string
	serviceName string
)

func stringFlagWithBind(envName, defaultValue, usage string) {
	flagName := strings.ReplaceAll(envName, "_", "-")
	rootCmd.PersistentFlags().String(flagName, defaultValue, usage)
	viper.BindPFlag(envName, rootCmd.PersistentFlags().Lookup(flagName))
}

func boolFlagWithBind(envName string, defaultValue bool, usage string) {
	flagName := strings.ReplaceAll(envName, "_", "-")
	rootCmd.PersistentFlags().Bool(flagName, defaultValue, usage)
	viper.BindPFlag(envName, rootCmd.PersistentFlags().Lookup(flagName))
}

func parseSystemContext(cmd *cobra.Command) (*types.SystemContext, error) {
	ctx, err := parse.SystemContextFromOptions(cmd)
	if err != nil {
		return nil, err
	}

	if config.ImageRegistryUsername != "" {
		ctx.DockerAuthConfig = &types.DockerAuthConfig{
			Username: config.ImageRegistryUsername,
			Password: config.ImageRegistryPassword,
		}
	}

	ctx.DockerInsecureSkipTLSVerify = types.NewOptionalBool(!config.ImageRegistryTLSVerify)
	ctx.OCIInsecureSkipTLSVerify = !config.ImageRegistryTLSVerify
	ctx.DockerDaemonInsecureSkipTLSVerify = !config.ImageRegistryTLSVerify

	return ctx, nil
}

func buildFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&imageName, "imageName", "i", "", "Image name")
	cmd.Flags().StringVarP(&imageTag, "imageTag", "t", "", "Image tag")
	cmd.Flags().StringVarP(&serviceName, "serviceName", "s", "", "KNative service name")
}
