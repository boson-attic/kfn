package dsl

import "github.com/slinkydeveloper/kfn/pkg/dsl/component"

func CreateAllResources(parsedSymbols map[string]component.Component) ([]interface{}, error) {
	resources := make([]interface{}, 0)
	for _, c := range parsedSymbols {
		newResources, err := c.GenerateDeployResources()
		if err != nil {
			return nil, err
		}
		if newResources != nil {
			resources = append(resources, newResources...)
		}
	}
	return resources, nil
}
