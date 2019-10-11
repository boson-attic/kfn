package dsl

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/slinkydeveloper/kfn/pkg/dsl/component"
	"github.com/slinkydeveloper/kfn/pkg/dsl/gen"
)

type ParseWires struct {
	*gen.BaseKfnListener

	parsedSymbolTable            map[string]component.Component
	componentsToAnonymousSymbols map[gen.IComponentContext]string

	collectedWires  [][]string
	collectedErrors []error
}

func NewParseWires(parsedSymbolTable map[string]component.Component, componentsToAnonymousSymbols map[gen.IComponentContext]string) *ParseWires {
	return &ParseWires{
		parsedSymbolTable:            parsedSymbolTable,
		componentsToAnonymousSymbols: componentsToAnonymousSymbols,
		collectedErrors:              make([]error, 0),
	}
}

func (ra *ParseWires) ExitWire(ctx *gen.WireContext) {
	wire := make([]string, 0)
	stmt := ctx.WireStmt().(*gen.WireStmtContext)
	for {
		leftWireElement := stmt.WireElement(0).(*gen.WireElementContext)
		leftWireComponent := ra.resolveComponent(leftWireElement)
		wire = append(wire, leftWireComponent)
		if stmt.WireStmt() != nil {
			stmt = stmt.WireStmt().(*gen.WireStmtContext)
		} else {
			rightWireElement := stmt.WireElement(1).(*gen.WireElementContext)
			rightWireComponent := ra.resolveComponent(rightWireElement)
			wire = append(wire, rightWireComponent)
			break
		}
	}
	ra.collectedWires = append(ra.collectedWires, wire)
}

func (ra *ParseWires) resolveComponent(wireElement *gen.WireElementContext) string {
	if wireElement.Ident() != nil {
		return wireElement.Ident().GetText()
	} else {
		return ra.componentsToAnonymousSymbols[wireElement.Component()]
	}
}

func (ra *ParseWires) Run(tree gen.IKfnContext) ([][]string, []error) {
	antlr.ParseTreeWalkerDefault.Walk(ra, tree)
	if len(ra.collectedErrors) != 0 {
		return nil, ra.collectedErrors
	}
	return ra.collectedWires, nil
}
