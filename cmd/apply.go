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
	"bufio"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/containers/buildah/pkg/unshare"
	"github.com/go-yaml/yaml"
	log "github.com/sirupsen/logrus"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/dsl"
	"github.com/slinkydeveloper/kfn/pkg/dsl/gen"
	"github.com/spf13/cobra"
	"os"
)

//TODO For now this command outputs yaml, in future it should apply it for real

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply [input_descriptor] [output_yaml]",
	Short: "Apply kfn build descriptor",
	RunE:  applyCmdFn,
	Args:  cobra.ExactArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		unshare.MaybeReexecUsingUserNamespace(false) // Do crazy stuff that allows buildah to work
		config.InitRunVariables()
		return config.InitBuildVariables(cmd)
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	buildFlags(applyCmd)
	runFlags(applyCmd)
}

func applyCmdFn(cmd *cobra.Command, args []string) error {
	// Parse the input
	input, err := antlr.NewFileStream(args[0])
	if err != nil {
		return err
	}

	lexer := gen.NewKfnLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := gen.NewKfnParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Kfn()

	log.Infof("Phase 1 - Create symbol table")

	phase1 := dsl.NewCreateDeclSymbolTablePhase()
	tree, unparsedSymbols, errs := phase1.Run(tree)
	if errs != nil {
		return errs[0]
	}
	log.Debugf("Unparsed symbols: %v", unparsedSymbols)

	log.Infof("Phase 2 - Remove anonymous symbols")

	phase2 := dsl.NewRemoveAnonymousSymbols(unparsedSymbols)
	tree, unparsedSymbols, componentsToAnonymousSymbols, errs := phase2.Run(tree)
	if errs != nil {
		return errs[0]
	}
	log.Debugf("Unparsed symbols: %v", unparsedSymbols)

	log.Infof("Phase 3 - Parse symbol values")

	parsedSymbols, err := dsl.ParseSymbolValues(unparsedSymbols)
	if err != nil {
		return err
	}
	log.Debugf("Parsed symbols: %v", parsedSymbols)

	log.Infof("Phase 4 - Parse wires")

	phase4 := dsl.NewParseWires(parsedSymbols, componentsToAnonymousSymbols)
	wires, errs := phase4.Run(tree)
	if errs != nil {
		return errs[0]
	}
	for _, w := range wires {
		log.Debugf("Wire: %+v", w)
	}

	log.Infof("Phase 5 - Check and expand wires")

	expandedWires, err := dsl.CheckAndExpandWires(wires, parsedSymbols)
	if err != nil {
		return err
	}
	for _, w := range expandedWires {
		log.Debugf("Expandend wire: %v", w)
	}

	log.Infof("Phase 6 - Build all functions")
	err = dsl.BuildAllFunctionImages(parsedSymbols)
	if err != nil {
		return err
	}

	log.Infof("Phase 7 - Create resource descriptors")
	log.Infof("Using namespace %s", config.Namespace)

	descriptors, err := dsl.CreateAllResources(parsedSymbols)
	if err != nil {
		return err
	}
	connectionsDescriptors, err := dsl.CreateAllWireResources(expandedWires, parsedSymbols)
	if err != nil {
		return err
	}

	log.Infof("All Resources created, writing to %s", args[1])

	return outputToYaml(append(descriptors, connectionsDescriptors...), args[1])
}

func outputToYaml(yamls []interface{}, outputFile string) error {
	outFile, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer outFile.Close()

	w := bufio.NewWriter(outFile)
	defer w.Flush()

	for i, y := range yamls {
		b, err := yaml.Marshal(y)
		if err != nil {
			return err
		}

		_, err = w.Write(b)
		if err != nil {
			return err
		}

		if i != len(yamls)-1 {
			_, err = w.WriteString("\n---\n\n")
			if err != nil {
				return err
			}
		}

		err = w.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}
