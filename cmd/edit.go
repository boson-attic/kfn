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
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/editors"
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"github.com/slinkydeveloper/kfn/pkg/util"
	"github.com/spf13/cobra"
	"path"
	"path/filepath"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit <function> <editor>",
	Args:  cobra.ExactArgs(2),
	Short: "Open the editor of your choice. Supported editors: idea, vscode",
	RunE:  editCmdFn,
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func editCmdFn(cmd *cobra.Command, args []string) error {
	functionPath := args[0]
	functionPath, err := filepath.Abs(functionPath)
	if err != nil {
		return err
	}

	editingDir := config.GetEditingDir(functionPath)

	language := languages.GetLanguage(path.Ext(functionPath))
	if language == languages.Unknown {
		return fmt.Errorf("Unknown language for function %s", functionPath)
	}

	ed := editors.GetEditor(args[1])
	if ed == editors.Unknown {
		return fmt.Errorf("Cannot recognize editor %s", args[1])
	}

	err = util.MkdirpIfNotExists(editingDir)
	if err != nil {
		return err
	}

	directory, descriptorFilename, err := language.ConfigureEditingDirectory(functionPath, editingDir)
	if err != nil {
		return err
	}

	err = ed.OpenEditor(directory, descriptorFilename)
	if err != nil {
		return err
	}

	return nil
}
