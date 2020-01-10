package dsl

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/slinkydeveloper/kfn/pkg/dsl/gen"
)

type CreateDeclSymbolTablePhase struct {
	*gen.BaseKfnListener

	unparsedSymbolTable map[string]gen.IComponentContext
	collectedErrors     []error
}

func NewCreateDeclSymbolTablePhase() *CreateDeclSymbolTablePhase {
	return &CreateDeclSymbolTablePhase{
		unparsedSymbolTable: make(map[string]gen.IComponentContext),
		collectedErrors:     make([]error, 0),
	}
}

func (cd *CreateDeclSymbolTablePhase) ExitDecl(c *gen.DeclContext) {
	ident := c.Ident().GetText()
	_, ok := cd.unparsedSymbolTable[ident]
	if ok {
		cd.collectedErrors = append(cd.collectedErrors, fmt.Errorf("Duplicated declaration '%s' at %d:%d", ident, c.GetStart().GetLine(), c.GetStart().GetColumn()))
	} else {
		value := c.Component()
		cd.unparsedSymbolTable[ident] = value
	}
}

func (cd *CreateDeclSymbolTablePhase) Run(tree gen.IKfnContext) (gen.IKfnContext, map[string]gen.IComponentContext, []error) {
	antlr.ParseTreeWalkerDefault.Walk(cd, tree)
	if len(cd.collectedErrors) != 0 {
		return nil, nil, cd.collectedErrors
	}
	return tree, cd.unparsedSymbolTable, nil
}
