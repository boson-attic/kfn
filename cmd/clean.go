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
	"github.com/slinkydeveloper/kfn/pkg/util"
	"github.com/spf13/cobra"
	"path/filepath"
)

var (
	global bool
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean <function_name>",
	Short: "Clean runtime & target directory for the provided function",
	RunE: func(cmd *cobra.Command, args []string) error {
		if global {
			return util.RmR(config.GetKfnDir())
		}
		functionPath := args[0]
		functionPath, err := filepath.Abs(functionPath)
		if err != nil {
			return err
		}

		targetDir := config.GetTargetDir(functionPath)
		editingDir := config.GetEditingDir(functionPath)

		return util.RmR(targetDir, editingDir)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().BoolVarP(&global, "global", "g", false, "Clean global kfn directory (default false)")
}
