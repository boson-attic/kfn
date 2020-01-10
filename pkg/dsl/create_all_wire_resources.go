package dsl

import (
	log "github.com/sirupsen/logrus"
	"github.com/slinkydeveloper/kfn/pkg/dsl/component"
)

func CreateAllWireResources(wires [][]string, parsedSymbols map[string]component.Component) ([]interface{}, error) {
	resources := make([]interface{}, 0)
	for _, w := range wires {
		newResources, err := createWireResources(w, parsedSymbols)
		if err != nil {
			return nil, err
		}
		if newResources != nil {
			resources = append(resources, newResources...)
		}
	}
	return resources, nil
}

func createWireResources(wire []string, parsedSymbols map[string]component.Component) ([]interface{}, error) {
	resources := make([]interface{}, 0)
	for i, c := range wire {
		var prev component.Component
		if i-1 >= 0 {
			prev = parsedSymbols[wire[i-1]]
		}
		var next component.Component
		if i+1 < len(wire) {
			next = parsedSymbols[wire[i+1]]
		}

		log.Debugf("Generating wire connection %v -> %v -> %v", prev, parsedSymbols[c], next)
		newResources, err := parsedSymbols[c].GenerateWireConnectionResources(prev, next)
		if err != nil {
			return nil, err
		}
		if newResources != nil {
			resources = append(resources, newResources...)
		}
	}
	return resources, nil
}
