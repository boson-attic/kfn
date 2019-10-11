package dsl

import "github.com/slinkydeveloper/kfn/pkg/dsl/component"

func BuildAllFunctionImages(parsedSymbols map[string]component.Component) error {
	for k, c := range parsedSymbols {
		switch c.(type) {
		case *component.Function:
			fn := c.(*component.Function)
			err := fn.Build()
			if err != nil {
				return err
			}
			parsedSymbols[k] = fn
		}
	}
	return nil
}
