package dsl

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/google/uuid"
	"github.com/slinkydeveloper/kfn/pkg/dsl/gen"
)

type RemoveAnonymousSymbols struct {
	*gen.BaseKfnListener

	unparsedSymbolTable          map[string]gen.IComponentContext
	componentsToAnonymousSymbols map[gen.IComponentContext]string
	collectedErrors              []error
}

func NewRemoveAnonymousSymbols(unparsedSymbolTable map[string]gen.IComponentContext) *RemoveAnonymousSymbols {
	return &RemoveAnonymousSymbols{
		unparsedSymbolTable:          unparsedSymbolTable,
		componentsToAnonymousSymbols: make(map[gen.IComponentContext]string),
		collectedErrors:              make([]error, 0),
	}
}

func (ra *RemoveAnonymousSymbols) ExitWireElement(ctx *gen.WireElementContext) {
	anonymousComponent := ctx.Component()
	if anonymousComponent != nil {
		key := uuid.New().String()
		ra.unparsedSymbolTable[key] = anonymousComponent
		ra.componentsToAnonymousSymbols[anonymousComponent] = key
	}
}

func (ra *RemoveAnonymousSymbols) Run(tree gen.IKfnContext) (gen.IKfnContext, map[string]gen.IComponentContext, map[gen.IComponentContext]string, []error) {
	antlr.ParseTreeWalkerDefault.Walk(ra, tree)
	if len(ra.collectedErrors) != 0 {
		return nil, nil, nil, ra.collectedErrors
	}
	return tree, ra.unparsedSymbolTable, ra.componentsToAnonymousSymbols, nil
}
