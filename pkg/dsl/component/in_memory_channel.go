package component

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/util"
)

type inMemoryChannel struct {
	name    string
	options map[string]string
}

func NewInMemoryChannel(name string, options map[string]string) Component {
	if name == "" {
		name = fmt.Sprintf("in-memory-ch-kfn-generated-%d", anonymousCounter)
		anonymousCounter++
	}
	return &inMemoryChannel{name, options}
}

func (k *inMemoryChannel) K8sName() string {
	return k.name
}

func (k *inMemoryChannel) K8sApiGroup() string {
	return MESSAGING_V1ALPHA1_API_GROUP
}

func (k *inMemoryChannel) K8sKind() string {
	return "InMemoryChannel"
}

func (k *inMemoryChannel) ComponentType() ComponentType {
	return Channel
}

func (k *inMemoryChannel) Validate() error {
	return nil
}

func (k *inMemoryChannel) Expand(component Component) Component {
	return nil
}

func (k *inMemoryChannel) CanConnectTo(component Component) bool {
	return util.AnyOf(component.ComponentType(), Channel, Service)
}

func (k *inMemoryChannel) IsValidWireStart() bool {
	return true
}

func (k *inMemoryChannel) GenerateDeployResources() ([]interface{}, error) {
	kch := map[string]interface{}{
		"apiVersion": k.K8sApiGroup(),
		"kind":       k.K8sKind(),
		"metadata": map[string]interface{}{
			"name":      k.name,
			"namespace": config.Namespace,
		},
	}
	return []interface{}{kch}, nil
}

func (k *inMemoryChannel) GenerateWireConnectionResources(previous Component, next Component) ([]interface{}, error) {
	if previous != nil && previous.ComponentType() == Channel {
		return []interface{}{generateChannelToChannelSub(previous, k, nil)}, nil
	}
	return []interface{}{}, nil
}

func (k *inMemoryChannel) String() string {
	return fmt.Sprintf("In Memory channel '%s'", k.name)
}
