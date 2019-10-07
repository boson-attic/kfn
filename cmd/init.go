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
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"github.com/spf13/cobra"
)

// initCmd represents the init parent command
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Init your function",
}

func init() {
	InitCmd.AddCommand(newInitCmd("js", "Javascript", languages.Javascript))
	InitCmd.AddCommand(newInitCmd("rust", "Rust", languages.Rust))
	rootCmd.AddCommand(InitCmd)
}

func newInitCmd(languageCmdName string, languageLongName string, language languages.Language) *cobra.Command {
	return &cobra.Command{
		Use:   fmt.Sprintf("%s [function_name] [directory]", languageCmdName),
		Args:  cobra.MaximumNArgs(2),
		Short: fmt.Sprintf("Bootstrap a %s function", languageLongName),
		RunE:  newInitCmdFn(language),
	}
}

func newInitCmdFn(language languages.Language) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return language.Bootstrap(resolveFunctionAndDir(args))
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
