package component

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/util"
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

func (k *kafkaChannel) K8sName() string {
	return k.name
}

func (k *kafkaChannel) K8sApiGroup() string {
	return MESSAGING_V1ALPHA1_API_GROUP
}

func (k *kafkaChannel) K8sKind() string {
	return "KafkaChannel"
}

func (k *kafkaChannel) ComponentType() ComponentType {
	return Channel
}

func (k *kafkaChannel) Validate() error {
	return nil
}

func (k *kafkaChannel) Expand(component Component) Component {
	return nil
}

func (k *kafkaChannel) CanConnectTo(component Component) bool {
	return util.AnyOf(component.ComponentType(), Channel, Service)
}

func (k *kafkaChannel) IsValidWireStart() bool {
	return true
}

func (k *kafkaChannel) GenerateDeployResources() ([]interface{}, error) {
	kch := map[string]interface{}{
		"apiVersion": k.K8sApiGroup(),
		"kind":       k.K8sKind(),
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

func (k *kafkaChannel) GenerateWireConnectionResources(previous Component, next Component) ([]interface{}, error) {
	if previous != nil && previous.ComponentType() == Channel {
		return []interface{}{generateChannelToChannelSub(previous, k, nil)}, nil
	}
	return []interface{}{}, nil
}

func (k *kafkaChannel) String() string {
	return fmt.Sprintf("Kafka channel '%s'", k.name)
}
