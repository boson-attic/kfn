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
	"fmt"
	"github.com/containers/buildah/pkg/unshare"
	log "github.com/sirupsen/logrus"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/spf13/cobra"
	serving "knative.dev/serving/pkg/client/clientset/versioned"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run <function_name>",
	Short: "Run the provided function",
	Long:  `TODO`,
	Args:  cobra.ExactArgs(1),
	Run:   runCmdFn,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		unshare.MaybeReexecUsingUserNamespace(false) // Do crazy stuff that allows buildah to work
		config.InitRunVariables()
		return config.InitBuildVariables(cmd)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	buildFlags(runCmd)
}

func runCmdFn(cmd *cobra.Command, args []string) {
	functionImage := buildCmdFn(cmd, args)

	log.Infof("Image %s pushed", functionImage.ImageName)

	kconfig, err := config.CreateK8sClientConfig()
	if err != nil {
		panic(fmt.Sprintf("Cannot create a k8s client config: %+v", err))
	}

	// create the clientset for k8s
	servingClient, err := serving.NewForConfig(kconfig)
	if err != nil {
		panic(fmt.Sprintf("Cannot create a serving client: %+v", err))
	}

	err = functionImage.RunImage(servingClient.ServingV1alpha1(), serviceName, config.NAMESPACE)

	if err != nil {
		panic(fmt.Sprintf("Cannot deploy the service: %+v", err))
	}

	log.Infof("Service %s deployed", serviceName)
}
