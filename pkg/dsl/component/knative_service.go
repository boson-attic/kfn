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
	case *inMemoryChannel:
		if next != nil {
			switch next.(type) {
			case *kafkaChannel:
				return []interface{}{k.generateChannelToChannelSub(previous.K8sName(), "InMemoryChannel", next.K8sName(), "KafkaChannel")}, nil
			case *inMemoryChannel:
				return []interface{}{k.generateChannelToChannelSub(previous.K8sName(), "InMemoryChannel", next.K8sName(), "InMemoryChannel")}, nil
			}
		} else {
			return []interface{}{k.generateChannelToChannelSub(previous.K8sName(), "InMemoryChannel", "", "")}, nil
		}
	case *kafkaChannel:
		if next != nil {
			switch next.(type) {
			case *kafkaChannel:
				return []interface{}{k.generateChannelToChannelSub(previous.K8sName(), "KafkaChannel", next.K8sName(), "KafkaChannel")}, nil
			case *inMemoryChannel:
				return []interface{}{k.generateChannelToChannelSub(previous.K8sName(), "KafkaChannel", next.K8sName(), "InMemoryChannel")}, nil
			}
		} else {
			return []interface{}{k.generateChannelToChannelSub(previous.K8sName(), "KafkaChannel", "", "")}, nil
		}
	}
	return []interface{}{}, nil
}

func (f *knativeService) generateChannelToChannelSub(previousChannelName string, previousChannelType string, nextChannelName string, nextChannelType string) map[string]interface{} {
	specMap := map[string]interface{}{
		"channel": map[string]interface{}{
			"apiVersion": "messaging.knative.dev/v1alpha1",
			"kind":       previousChannelType,
			"name":       previousChannelName,
		},
		"subscriber": map[string]interface{}{
			"ref": map[string]interface{}{
				"apiVersion": "serving.knative.dev/v1alpha1",
				"kind":       "Service",
				"name":       f.serviceName,
			},
		},
	}

	if nextChannelName != "" {
		specMap["reply"] = map[string]interface{}{
			"channel": map[string]interface{}{
				"apiVersion": "messaging.knative.dev/v1alpha1",
				"kind":       nextChannelType,
				"name":       nextChannelName,
			},
		}
	}

	var subName string
	if nextChannelName != "" {
		subName = fmt.Sprintf("%s-%s-%s", previousChannelName, f.serviceName, nextChannelName)
	} else {
		subName = fmt.Sprintf("%s-%s", previousChannelName, f.serviceName)
	}

	return map[string]interface{}{
		"apiVersion": "messaging.knative.dev/v1alpha1",
		"kind":       "Subscription",
		"metadata": map[string]interface{}{
			"name":      subName,
			"namespace": config.Namespace,
		},
		"spec": specMap,
	}
}

func (k *knativeService) String() string {
	return fmt.Sprintf("KnativeService '%s'", k.serviceName)
}
