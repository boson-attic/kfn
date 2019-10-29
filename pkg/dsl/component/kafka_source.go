package component

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/util"
)

const (
	CONSUMER_GROUP    = "consumer_group"
	TOPICS            = "topics"
	BOOTSTRAP_SERVERS = "bootstrap_servers"
)

type kafkaSource struct {
	name                     string
	consumerGroup            string
	bootstrapServers         string
	topics                   string
	componentOutboundChannel Component
	options                  map[string]string
}

func NewKafkaSource(name string, options map[string]string) Component {
	if name == "" {
		name = fmt.Sprintf("cron-source-kfn-generated-%d", anonymousCounter)
		anonymousCounter++
	}
	return &kafkaSource{name: name, options: options}
}

func (k *kafkaSource) K8sName() string {
	return k.name
}

func (k *kafkaSource) K8sApiGroup() string {
	return SOURCES_v1ALPHA1_API_GROUP
}

func (k *kafkaSource) K8sKind() string {
	return "KafkaSource"
}

func (k *kafkaSource) ComponentType() ComponentType {
	return Source
}

func (k *kafkaSource) Validate() error {
	if val, ok := k.options[CONSUMER_GROUP]; !ok {
		k.consumerGroup = "knative-group"
	} else {
		k.consumerGroup = val
	}

	if val, ok := k.options[BOOTSTRAP_SERVERS]; ok {
		k.bootstrapServers = val
	} else {
		return fmt.Errorf("Missing %s option", BOOTSTRAP_SERVERS)
	}

	if val, ok := k.options[TOPICS]; ok {
		k.topics = val
	} else {
		return fmt.Errorf("Missing %s option", TOPICS)
	}

	k.componentOutboundChannel = defaultExpansionChannelFactory(fmt.Sprintf("%s-outbound-ch", k.name), nil)

	return nil
}

func (k *kafkaSource) Expand(component Component) Component {
	if component.ComponentType() == Service {
		return defaultExpansionChannelFactory("", nil)
	}
	return nil
}

func (k *kafkaSource) CanConnectTo(component Component) bool {
	return util.AnyOf(component.ComponentType(), Channel, Service)
}

func (k *kafkaSource) IsValidWireStart() bool {
	return true
}

func (k *kafkaSource) GenerateDeployResources() ([]interface{}, error) {
	ch, err := k.componentOutboundChannel.GenerateDeployResources()
	if err != nil {
		return nil, err
	}

	s := map[string]interface{}{
		"apiVersion": k.K8sApiGroup(),
		"kind":       k.K8sKind(),
		"metadata": map[string]interface{}{
			"name":      fmt.Sprintf("%s", k.name),
			"namespace": config.Namespace,
		},
		"spec": map[string]interface{}{
			"consumerGroup":    k.consumerGroup,
			"bootstrapServers": k.bootstrapServers,
			"topics":           k.topics,
			"sink":             generateRef(k.componentOutboundChannel),
		},
	}

	return append(ch, s), nil
}

func (k *kafkaSource) GenerateWireConnectionResources(previous Component, next Component) ([]interface{}, error) {
	// Next Resource is a channel by expansion rules, previous is nil
	return []interface{}{generateChannelToChannelSub(k.componentOutboundChannel, k, next)}, nil
}

func (k *kafkaSource) String() string {
	return fmt.Sprintf("Cron source '%s'", k.name)
}
