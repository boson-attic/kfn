/*
Copyright © Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/containers/buildah/pkg/unshare"
	"github.com/mitchellh/go-homedir"
	"github.com/slinkydeveloper/kfn/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var cfgFile string
var k8sNamespace string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kfn",
	Short: "TODO",
	Long:  `TODO`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		unshare.MaybeReexecUsingUserNamespace(false) // Do crazy stuff that allows buildah to work
		pkg.InitVariables()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kfn.yaml or $(pwd)/.kfn.yaml")

	rootCmd.PersistentFlags().BoolP(pkg.VERBOSE, "v", false, "verbose output")
	viper.BindPFlag(pkg.VERBOSE, rootCmd.PersistentFlags().Lookup(pkg.VERBOSE))

	stringFlagWithBind(pkg.REGISTRY, "", "Docker registry where to push the image")
	stringFlagWithBind(pkg.REGISTRY_USERNAME, "", "Username to access docker registry")
	stringFlagWithBind(pkg.REGISTRY_PASSWORD, "", "Password to access docker registry")
	stringFlagWithBind(pkg.KUBECONFIG, "", "kubeconfig")

	rootCmd.PersistentFlags().StringVarP(&k8sNamespace, "namespace", "n", "default", "K8s namespace where to run the service")
}

func stringFlagWithBind(envName, defaultValue, usage string) {
	flagName := strings.ReplaceAll(envName, "_", "-")
	rootCmd.PersistentFlags().String(flagName, defaultValue, usage)
	viper.BindPFlag(envName, rootCmd.PersistentFlags().Lookup(flagName))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kfn" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".kfn")
	}

	viper.AutomaticEnv()

	_ = viper.ReadInConfig()
}
