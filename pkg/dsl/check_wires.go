package dsl

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/slinkydeveloper/kfn/pkg/dsl/component"
)

func CheckAndExpandWires(wires [][]string, symbolsTable map[string]component.Component) ([][]string, error) {
	result := make([][]string, len(wires))
	for i, w := range wires {
		var err error
		result[i], err = checkAndExpand(w, symbolsTable)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func checkAndExpand(wire []string, symbolsTable map[string]component.Component) ([]string, error) {
	result := make([]string, 0)
	result = append(result, wire[0])
	if !symbolsTable[wire[0]].IsValidWireStart() {
		return nil, fmt.Errorf("component %v is not a valid wire beginning", symbolsTable[wire[0]])
	}
	for i := 1; i < len(wire); i++ {
		prev := symbolsTable[wire[i-1]]
		actual := symbolsTable[wire[i]]
		if !prev.CanConnectTo(actual) {
			return nil, fmt.Errorf("component %v can't connect to %v", prev, actual)
		}
		newComponent := prev.Expand(actual)
		if newComponent != nil {
			key := uuid.New().String()
			symbolsTable[key] = newComponent
			result = append(result, key, wire[i])
		} else {
			result = append(result, wire[i])
		}
	}
	return result, nil
}
