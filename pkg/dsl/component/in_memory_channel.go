package component

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg/config"
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

func (k inMemoryChannel) K8sName() string {
	return k.name
}

func (k *inMemoryChannel) Validate() error {
	return nil
}

func (k inMemoryChannel) Expand(component Component) Component {
	return nil
}

func (k inMemoryChannel) CanConnectTo(component Component) bool {
	switch component.(type) {
	case *Function:
		return true
	case *inMemoryChannel:
		return true
	case *kafkaChannel:
		return true
	case *knativeService:
		return true
	}
	return false
}

func (k inMemoryChannel) IsValidWireStart() bool {
	return true
}

func (k inMemoryChannel) GenerateDeployResources() ([]interface{}, error) {
	kch := map[string]interface{}{
		"apiVersion": "messaging.knative.dev/v1alpha1",
		"kind":       "InMemoryChannel",
		"metadata": map[string]interface{}{
			"name":      k.name,
			"namespace": config.Namespace,
		},
	}
	return []interface{}{kch}, nil
}

func (k inMemoryChannel) GenerateWireConnectionResources(previous Component, next Component) ([]interface{}, error) {
	if previous != nil {
		switch previous.(type) {
		case *inMemoryChannel:
			return []interface{}{k.generateIMChannelSub(previous.(*inMemoryChannel))}, nil
		case *kafkaChannel:
			return []interface{}{k.generateKafkaChannelSub(previous.(*kafkaChannel))}, nil
		}
	}
	return []interface{}{}, nil
}

func (k inMemoryChannel) generateKafkaChannelSub(previousChannel *kafkaChannel) map[string]interface{} {
	return map[string]interface{}{
		"apiVersion": "messaging.knative.dev/v1alpha1",
		"kind":       "Subscription",
		"metadata": map[string]interface{}{
			"name":      fmt.Sprintf("%s-%s", previousChannel.name, k.name),
			"namespace": config.Namespace,
		},
		"spec": map[string]interface{}{
			"channel": map[string]interface{}{
				"apiVersion": "messaging.knative.dev/v1alpha1",
				"kind":       "KafkaChannel",
				"name":       previousChannel.name,
			},
			"subscriber": map[string]interface{}{
				"ref": map[string]interface{}{
					"apiVersion": "messaging.knative.dev/v1alpha1",
					"kind":       "InMemoryChannel",
					"name":       k.name,
				},
			},
		},
	}
}

func (k inMemoryChannel) generateIMChannelSub(previousChannel *inMemoryChannel) map[string]interface{} {
	return map[string]interface{}{
		"apiVersion": "messaging.knative.dev/v1alpha1",
		"kind":       "Subscription",
		"metadata": map[string]interface{}{
			"name":      fmt.Sprintf("%s-%s", previousChannel.name, k.name),
			"namespace": config.Namespace,
		},
		"spec": map[string]interface{}{
			"channel": map[string]interface{}{
				"apiVersion": "messaging.knative.dev/v1alpha1",
				"kind":       "InMemoryChannel",
				"name":       previousChannel.name,
			},
			"subscriber": map[string]interface{}{
				"ref": map[string]interface{}{
					"apiVersion": "messaging.knative.dev/v1alpha1",
					"kind":       "InMemoryChannel",
					"name":       k.name,
				},
			},
		},
	}
}

func (k inMemoryChannel) String() string {
	return fmt.Sprintf("In Memory channel '%s'", k.name)
}
