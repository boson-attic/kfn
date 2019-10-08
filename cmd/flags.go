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
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

var (
	imageName string
	imageTag string
	serviceName string
)

func stringFlagWithBind(flagSet *pflag.FlagSet, envName, shorthandFlag, defaultValue, usage string) {
	flagName := strings.ReplaceAll(envName, "_", "-")
	flagSet.StringP(flagName, shorthandFlag, defaultValue, usage)
	_ = viper.BindPFlag(envName, flagSet.Lookup(flagName))
}

func boolFlagWithBind(flagSet *pflag.FlagSet, envName string, shorthandFlag string, defaultValue bool, usage string) {
	flagName := strings.ReplaceAll(envName, "_", "-")
	flagSet.BoolP(flagName, shorthandFlag, defaultValue, usage)
	_ = viper.BindPFlag(envName, flagSet.Lookup(flagName))
}

func buildFlags(cmd *cobra.Command) {
	stringFlagWithBind(cmd.Flags(), config.REGISTRY, "", "", "Docker registry where to push the image")
	stringFlagWithBind(cmd.Flags(), config.REGISTRY_USERNAME, "", "", "Username to access docker registry")
	stringFlagWithBind(cmd.Flags(), config.REGISTRY_PASSWORD, "", "", "Password to access docker registry")
	boolFlagWithBind(cmd.Flags(), config.REGISTRY_TLS_VERIFY, "", true, "TLS Verify when accessing the docker registry")
	cmd.Flags().StringVarP(&imageName, "imageName", "i", "", "Image name")
	cmd.Flags().StringVarP(&imageTag, "imageTag", "t", "", "Image tag")
	cmd.Flags().StringVarP(&serviceName, "serviceName", "s", "", "KNative service name")
}

func runFlags(cmd *cobra.Command) {
	stringFlagWithBind(cmd.Flags(), config.KUBECONFIG, "", "", "Kubeconfig")
	stringFlagWithBind(cmd.Flags(), config.NAMESPACE, "", "default", "K8s namespace where to run the service")
}