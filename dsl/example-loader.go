package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/slinkydeveloper/kfn/pkg/dsl"
	"github.com/slinkydeveloper/kfn/pkg/dsl/gen"
	"os"
)

func Load(file string) error {
	input, err := antlr.NewFileStream(file)
	if err != nil {
		return err
	}

	lexer := gen.NewKfnLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := gen.NewKfnParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Kfn()

	phase1 := dsl.NewCreateDeclSymbolTablePhase()
	tree, unparsedSymbols, errs := phase1.Run(tree)
	if errs != nil {
		panic(fmt.Sprintf("%+v", errs))
	}
	fmt.Printf("%v\n", unparsedSymbols)

	phase2 := dsl.NewRemoveAnonymousSymbols(unparsedSymbols)
	tree, unparsedSymbols, componentsToAnonymousSymbols, errs := phase2.Run(tree)
	if errs != nil {
		panic(fmt.Sprintf("%+v", errs))
	}
	fmt.Printf("%v\n", unparsedSymbols)
	fmt.Printf("%v\n", componentsToAnonymousSymbols)

	parsedSymbols, err := dsl.ParseSymbolValues(unparsedSymbols)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", parsedSymbols)

	phase4 := dsl.NewParseWires(parsedSymbols, componentsToAnonymousSymbols)
	wires, errs := phase4.Run(tree)
	if errs != nil {
		panic(fmt.Sprintf("%+v", errs))
	}
	for _, w := range wires {
		fmt.Printf("Wire: %+v\n", w)
	}

	expandedWires, err := dsl.CheckAndExpandWires(wires, parsedSymbols)
	if err != nil {
		panic(err)
	}
	for _, w := range expandedWires {
		fmt.Printf("Expandend wire: %v\n", w)
	}

	return nil
}

func main() {
	err := Load(os.Args[1])
	if err != nil {
		panic(err)
	}
}
