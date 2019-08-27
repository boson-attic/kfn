/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"github.com/slinkydeveloper/kfn/pkg/kfn"
	"github.com/slinkydeveloper/kfn/pkg/kfn/image"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	serving "knative.dev/serving/pkg/client/clientset/versioned"
	"path"
	"strings"
)

var imageName string
var imageTag string
var serviceName string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [function_name]",
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
	runCmd.Flags().StringVarP(&imageTag, "imageTag", "t", "latest", "Image tag")
	runCmd.Flags().StringVarP(&serviceName, "serviceName", "s", "", "KNative service name")
}

func runCmdFn(cmd *cobra.Command, args []string) {
	function_path := args[0]

	logf("Loading function %v", function_path)

	err := kfn.LoadFunction(function_path)

	if err != nil {
		panic(fmt.Sprintf("Error while loading the function: %v", err))
	}

	logf("Function loaded")

	logf("Loading runtime")

	err = kfn.DownloadRuntimeAndCopyRequiredFiles()

	if err != nil {
		panic(fmt.Sprintf("Error while loading the runtime: %v", err))
	}

	logf("Runtime loaded")

	if len(imageName) == 0 {
		base := path.Base(function_path)
		imageName = strings.TrimSuffix(base, path.Ext(base))
	}

	if len(serviceName) == 0 {
		serviceName = imageName
	}

	functionImage := image.FunctionImage{
		ImageName:     imageName,
		ImageRegistry: DockerRegistry,
		Tag:           imageTag,
	}

	logf("Building image")

	imageId, err := functionImage.BuildImage(kfn.TargetDirectory)

	if err != nil {
		panic(fmt.Sprintf("Error while building the image: %v", err))
	}

	logf("Image built: %v", imageId)

	err = functionImage.PushImage(imageId)

	if err != nil {
		panic(fmt.Sprintf("Error while pushing the image: %v", err))
	}

	logf("Image pushed")

	config, err := clientcmd.BuildConfigFromFlags("", Kubeconfig)
	if err != nil {
		panic(fmt.Sprintf("Cannot create a k8s client config: %+v", err))
	}

	// create the clientset for k8s
	servingClient, err := serving.NewForConfig(config)
	if err != nil {
		panic(fmt.Sprintf("Cannot create a serving client: %+v", err))
	}

	err = functionImage.RunImage(servingClient.ServingV1alpha1(), serviceName, k8sNamespace)

	if err != nil {
		panic(fmt.Sprintf("Cannot create a serving client: %+v", err))
	}

	logf("KNative service deployed")

}

func logf(format string, values ...interface{}) {
	if Verbose {
		fmt.Printf(format+"\n", values...)
	}
}
