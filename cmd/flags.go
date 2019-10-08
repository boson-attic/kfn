package cmd

import (
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

func boolFlagWithBind(envName string, shorthandFlag string, defaultValue bool, usage string) {
	flagName := strings.ReplaceAll(envName, "_", "-")
	rootCmd.PersistentFlags().BoolP(flagName, shorthandFlag, defaultValue, usage)
	viper.BindPFlag(envName, rootCmd.PersistentFlags().Lookup(flagName))
}

func buildFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&imageName, "imageName", "i", "", "Image name")
	cmd.Flags().StringVarP(&imageTag, "imageTag", "t", "", "Image tag")
	cmd.Flags().StringVarP(&serviceName, "serviceName", "s", "", "KNative service name")
}
