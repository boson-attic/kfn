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
	"github.com/containers/buildah/pkg/parse"
	"github.com/containers/image/types"
	log "github.com/sirupsen/logrus"
	"github.com/slinkydeveloper/kfn/pkg"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"github.com/slinkydeveloper/kfn/pkg/util"
	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	serving "knative.dev/serving/pkg/client/clientset/versioned"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var imageName string
var imageTag string
var serviceName string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run <function_name>",
	Short: "Run the provided function",
	Long:  `TODO`,
	Args:  cobra.ExactArgs(1),
	Run:   runCmdFn,
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	runCmd.Flags().StringVarP(&imageName, "imageName", "i", "", "Image name")
	runCmd.Flags().StringVarP(&imageTag, "imageTag", "t", "", "Image tag")
	runCmd.Flags().StringVarP(&serviceName, "serviceName", "s", "", "KNative service name")
}

func runCmdFn(cmd *cobra.Command, args []string) {
	if err := util.RmR(config.TargetDir); err != nil {
		panic(err)
	}

	log.Infof("Using Kubeconfig: %v\n", config.Kubeconfig)
	log.Infof("Using Docker registry: %v\n", config.ImageRegistry)

	functionPath := args[0]
	functionPath, err := filepath.Abs(functionPath)
	if err != nil {
		panic(err)
	}

	language := languages.GetLanguage(path.Ext(functionPath))
	if language == languages.Unknown {
		panic(fmt.Sprintf("Unknown language for function %s", functionPath))
	}

	if len(imageName) == 0 {
		base := path.Base(functionPath)
		imageName = strings.TrimSuffix(base, path.Ext(base))
	}

	if len(serviceName) == 0 {
		serviceName = imageName
	}

	ctx, err := parseSystemContext(cmd)
	if err != nil {
		panic(fmt.Sprintf("Error while trying to infer context: %v", err))
	}

	functionImage, err := pkg.Build(functionPath, language, imageName, imageTag, ctx)
	if err != nil {
		panic(fmt.Sprintf("Error while building the image: %v", err))
	}

	log.Infof("Image %+v pushed", functionImage)

	var kconfig *rest.Config
	if os.Getenv("KFN_IN_CLUSTER") == "true" {
		kconfig, err = rest.InClusterConfig()
	} else {
		if config.Kubeconfig != "" {
			kconfig, err = clientcmd.BuildConfigFromFlags("", config.Kubeconfig)
		} else {
			kconfig, err = clientcmd.BuildConfigFromKubeconfigGetter("", clientcmd.NewDefaultClientConfigLoadingRules().Load)
		}
	}
	if err != nil {
		panic(fmt.Sprintf("Cannot create a k8s client config: %+v", err))
	}

	// create the clientset for k8s
	servingClient, err := serving.NewForConfig(kconfig)
	if err != nil {
		panic(fmt.Sprintf("Cannot create a serving client: %+v", err))
	}

	err = functionImage.RunImage(servingClient.ServingV1alpha1(), serviceName, k8sNamespace)

	if err != nil {
		panic(fmt.Sprintf("Cannot create a serving client: %+v", err))
	}

	log.Infof("Service %s deployed", serviceName)
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
