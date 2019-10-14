package component

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg/config"
)

type knativeService struct {
	options     map[string]string
	serviceName string
	image       string
}

func NewKnativeService(name string, options map[string]string) Component {
	if name == "" {
		name = fmt.Sprintf("ksvc-kfn-generated-%d", anonymousCounter)
		anonymousCounter++
	}
	return &knativeService{serviceName: name, options: options}
}

func (k *knativeService) K8sName() string {
	return k.serviceName
}

func (k *knativeService) Validate() error {
	if k.options["image"] == "" {
		return fmt.Errorf("you must specify the image for service '%s'", k.serviceName)
	}

	k.image = k.options["image"]

	return nil
}

func (k *knativeService) Expand(component Component) Component {
	switch component.(type) {
	case *Function:
		return defaultExpansionChannelFactory("", nil)
	case *knativeService:
		return defaultExpansionChannelFactory("", nil)
	}
	return nil
}

func (k *knativeService) CanConnectTo(component Component) bool {
	switch component.(type) {
	case *knativeService:
		return true
	case *kafkaChannel:
		return true
	case *Function:
		return true
	}
	return false
}

func (k *knativeService) IsValidWireStart() bool {
	return false
}

func (k *knativeService) GenerateDeployResources() ([]interface{}, error) {
	return []interface{}{k.generateService()}, nil
}

func (k *knativeService) generateService() interface{} {
	return map[string]interface{}{
		"apiVersion": "serving.knative.dev/v1alpha1",
		"kind":       "Service",
		"metadata": map[string]string{
			"name":      k.serviceName,
			"namespace": config.Namespace,
		},
		"spec": map[string]interface{}{
			"template": map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []interface{}{map[string]interface{}{
						"image": k.image,
					}},
				},
			},
		},
	}
}

func (k *knativeService) GenerateWireConnectionResources(previous Component, next Component) ([]interface{}, error) {
	switch previous.(type) {
	case *kafkaChannel:
		if next != nil {
			switch next.(type) {
			case *kafkaChannel:
				return []interface{}{k.generateChannelToChannelSub(previous.(*kafkaChannel), next.(*kafkaChannel))}, nil
			}
		} else {
			return []interface{}{k.generateChannelSub(previous.(*kafkaChannel))}, nil
		}
	}
	return []interface{}{}, nil
}

func (k *knativeService) generateChannelToChannelSub(previousChannel *kafkaChannel, nextChannel *kafkaChannel) map[string]interface{} {
	return map[string]interface{}{
		"apiVersion": "messaging.knative.dev/v1alpha1",
		"kind":       "Subscription",
		"metadata": map[string]interface{}{
			"name":      fmt.Sprintf("%s-%s-%s", previousChannel.name, k.serviceName, nextChannel.name),
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
					"apiVersion": "serving.knative.dev/v1alpha1",
					"kind":       "Service",
					"name":       k.serviceName,
				},
			},
			"reply": map[string]interface{}{
				"channel": map[string]interface{}{
					"apiVersion": "messaging.knative.dev/v1alpha1",
					"kind":       "KafkaChannel",
					"name":       nextChannel.name,
				},
			},
		},
	}
}

func (k *knativeService) generateChannelSub(previousChannel *kafkaChannel) map[string]interface{} {
	return map[string]interface{}{
		"apiVersion": "messaging.knative.dev/v1alpha1",
		"kind":       "Subscription",
		"metadata": map[string]interface{}{
			"name":      fmt.Sprintf("%s-%s", previousChannel.name, k.serviceName),
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
					"apiVersion": "serving.knative.dev/v1alpha1",
					"kind":       "Service",
					"name":       k.serviceName,
				},
			},
		},
	}
}

func (k *knativeService) String() string {
	return fmt.Sprintf("KnativeService '%s'", k.serviceName)
}
