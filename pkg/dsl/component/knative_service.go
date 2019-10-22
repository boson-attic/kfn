package component

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/util"
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

func (f *knativeService) K8sApiGroup() string {
	return SERVING_V1ALPHA1_API_GROUP
}

func (f *knativeService) K8sKind() string {
	return "Service"
}

func (f *knativeService) ComponentType() ComponentType {
	return Service
}

func (k *knativeService) Validate() error {
	if k.options["image"] == "" {
		return fmt.Errorf("you must specify the image for service '%s'", k.serviceName)
	}

	k.image = k.options["image"]

	return nil
}

func (k *knativeService) Expand(component Component) Component {
	if component.ComponentType() == Service {
		return defaultExpansionChannelFactory("", nil)
	}
	return nil
}

func (k *knativeService) CanConnectTo(component Component) bool {
	return util.AnyOf(component.ComponentType(), Channel, Service)
}

func (k *knativeService) IsValidWireStart() bool {
	return false
}

func (k *knativeService) GenerateDeployResources() ([]interface{}, error) {
	return []interface{}{k.generateService()}, nil
}

func (k *knativeService) generateService() interface{} {
	return map[string]interface{}{
		"apiVersion": k.K8sApiGroup(),
		"kind":       k.K8sKind(),
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
	// Previous can't be null because ksvc is not a valid wire start
	if previous.ComponentType() == Channel && (next == nil || next.ComponentType() == Channel) {
		return []interface{}{generateChannelToChannelSub(previous, k, next)}, nil
	}
	return []interface{}{}, nil
}

func (k *knativeService) String() string {
	return fmt.Sprintf("KnativeService '%s'", k.serviceName)
}
