/*
Copyright © 2019 Red Hat, Inc.

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
