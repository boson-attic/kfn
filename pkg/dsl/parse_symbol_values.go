package dsl

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/slinkydeveloper/kfn/pkg/dsl/component"
	"github.com/slinkydeveloper/kfn/pkg/dsl/gen"
)

func ParseSymbolValues(unparsedSymbolTable map[string]gen.IComponentContext) (map[string]component.Component, error) {
	result := make(map[string]component.Component)
	for k, v := range unparsedSymbolTable {
		var err error
		result[k], err = ParseComponent(v)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func ParseComponent(c gen.IComponentContext) (component.Component, error) {
	comp := c.(*gen.ComponentContext)
	componentIdent := comp.ComponentIdent()

	if componentIdent == nil {
		// No component ident means there is a string literal that means function
		val := unquote(comp.ComponentValue().GetText())
		function := component.NewFunction(val, nil)
		err := function.Validate()
		if err != nil {
			return function, nil
		} else {
			return nil, err
		}
	}

	componentType := componentIdent.GetText()
	componentFactory := component.ResolveComponentFactory(componentType)
	if componentFactory == nil {
		return nil, fmt.Errorf("Unknown component %s at %d:%d", componentType, componentIdent.GetStart().GetLine(), componentIdent.GetStart().GetColumn())
	}

	var componentValue string
	if comp.ComponentValue() != nil {
		componentValue = unquote(comp.ComponentValue().GetText())
	}

	options := make(map[string]string)
	componentOptionsList := comp.ComponentOptionList()
	if componentOptionsList != nil {
		collector := componentOptionCollector{options: options}
		antlr.ParseTreeWalkerDefault.Walk(&collector, componentOptionsList)
	}

	return componentFactory(componentValue, options), nil
}

type componentOptionCollector struct {
	*gen.BaseKfnListener

	options map[string]string
}

func (co componentOptionCollector) ExitComponentOption(c *gen.ComponentOptionContext) {
	key := c.ComponentOptionIdent().GetText()
	value := unquote(c.ComponentOptionValue().GetText())
	co.options[key] = value
}

//TODO must improve removing escaped codes
func unquote(stringLiteral string) string {
	return stringLiteral[1 : len(stringLiteral)-1]
}
