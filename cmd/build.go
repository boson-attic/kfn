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
	"github.com/slinkydeveloper/kfn/pkg"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/image"
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build <function_name>",
	Short: "Build the function image",
	Run: func(cmd *cobra.Command, args []string) {
		buildCmdFn(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildFlags(buildCmd)
}

func buildCmdFn(cmd *cobra.Command, args []string) image.FunctionImage {
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

	ctx, err := config.ParseSystemContext(cmd)
	if err != nil {
		panic(fmt.Sprintf("Error while trying to infer context: %v", err))
	}

	functionImage, err := pkg.Build(functionPath, language, imageName, imageTag, ctx)
	if err != nil {
		panic(fmt.Sprintf("Error while building the image: %v", err))
	}

	log.Infof("Image %+v pushed", functionImage)

	return functionImage
}
