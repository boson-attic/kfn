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
	"github.com/spf13/cobra"
)

// initCmd represents the init parent command
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Init your function",
}

func init() {
	InitCmd.AddCommand(newInitCmd("js", kfn.Javascript))
	rootCmd.AddCommand(InitCmd)
}

func newInitCmd(languageCmdName string, language kfn.Language) *cobra.Command {
	return &cobra.Command{
		Use:   fmt.Sprintf("%s [function_name] [directory]", languageCmdName),
		Args:  cobra.MaximumNArgs(2),
		Short: "Bootstrap a function",
		RunE:  newInitCmdFn(language),
	}
}

func newInitCmdFn(language kfn.Language) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		functionName, directory := resolveFunctionAndDir(args)

		bootstrapper := kfn.ResolveBootrapper(kfn.Javascript)

		return (*bootstrapper).Bootstrap(functionName, directory)
	}
}

func resolveFunctionAndDir(args []string) (string, string) {
	if len(args) == 2 {
		return args[0], args[1]
	} else if len(args) == 1 {
		return args[0], ""
	} else {
		return "example-fn", ""
	}
}
