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
	"os"

	"github.com/containers/buildah/pkg/unshare"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var Verbose bool
var k8sNamespace string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kfn",
	Short: "TODO",
	Long:  `TODO`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		unshare.MaybeReexecUsingUserNamespace(false) // Do crazy stuff that allows buildah to work
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
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().String("registry", "", "Docker registry where to push the image")
	viper.BindPFlag("registry", rootCmd.PersistentFlags().Lookup("registry"))
	rootCmd.PersistentFlags().String("registry-username", "", "Username to access docker registry")
	viper.BindPFlag("registry_username", rootCmd.PersistentFlags().Lookup("registry-username"))
	rootCmd.PersistentFlags().String("registry-password", "", "Password to access docker registry")
	viper.BindPFlag("registry_password", rootCmd.PersistentFlags().Lookup("registry-password"))

	rootCmd.PersistentFlags().String("kubeconfig", "", "kubeconfig")
	viper.BindPFlag("kubeconfig", rootCmd.PersistentFlags().Lookup("kubeconfig"))

	rootCmd.PersistentFlags().StringVarP(&k8sNamespace, "namespace", "n", "default", "K8s namespace where to run the service")
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

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
