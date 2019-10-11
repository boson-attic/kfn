package component

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg/config"
)

type kafkaChannel struct {
	name    string
	options map[string]string
}

func NewKafkaChannel(name string, options map[string]string) Component {
	if name == "" {
		name = fmt.Sprintf("kafka-ch-kfn-generated-%d", anonymousCounter)
		anonymousCounter++
	}
	return &kafkaChannel{name, options}
}

func (k kafkaChannel) Validate() error {
	return nil
}

func (k kafkaChannel) Expand(component Component) Component {
	return nil
}

func (k kafkaChannel) CanConnectTo(component Component) bool {
	switch component.(type) {
	case *Function:
		return true
	case *kafkaChannel:
		return true
	}
	return false
}

func (k kafkaChannel) IsValidWireStart() bool {
	return true
}

func (k kafkaChannel) GenerateDeployResources() ([]interface{}, error) {
	kch := map[string]interface{}{
		"apiVersion": "messaging.knative.dev/v1alpha1",
		"kind":       "KafkaChannel",
		"metadata": map[string]interface{}{
			"name":      k.name,
			"namespace": config.Namespace,
		},
		"spec": map[string]interface{}{
			"numPartitions":     10,
			"replicationFactor": 1,
		},
	}
	return []interface{}{kch}, nil
}

func (k kafkaChannel) GenerateWireConnectionResources(previous Component, next Component) ([]interface{}, error) {
	if previous != nil {
		switch previous.(type) {
		case *kafkaChannel:
			return []interface{}{k.generateChannelSub(previous.(*kafkaChannel))}, nil
		}
	}
	return []interface{}{}, nil
}

func (k kafkaChannel) generateChannelSub(previousChannel *kafkaChannel) map[string]interface{} {
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
					"kind":       "KafkaChannel",
					"name":       k.name,
				},
			},
		},
	}
}

func (k kafkaChannel) String() string {
	return fmt.Sprintf("Kafka channel '%s'", k.name)
}
