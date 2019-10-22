package component

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/util"
)

type cronSource struct {
	name    string
	options map[string]string
}

func NewCronSource(name string, options map[string]string) Component {
	if name == "" {
		name = fmt.Sprintf("cron-source-kfn-generated-%d", anonymousCounter)
		anonymousCounter++
	}
	return &cronSource{name, options}
}

func (k *cronSource) K8sName() string {
	return k.name
}

func (k *cronSource) K8sApiGroup() string {
	return SOURCES_v1ALPHA1_API_GROUP
}

func (k *cronSource) K8sKind() string {
	return "CronJobSource"
}

func (k *cronSource) ComponentType() ComponentType {
	return Source
}

func (k *cronSource) Validate() error {
	return nil
}

func (k *cronSource) Expand(component Component) Component {
	if component.ComponentType() == Service {
		return defaultExpansionChannelFactory("", nil)
	}
	return nil
}

func (k *cronSource) CanConnectTo(component Component) bool {
	return util.AnyOf(component.ComponentType(), Channel, Service)
}

func (k *cronSource) IsValidWireStart() bool {
	return true
}

func (k *cronSource) GenerateDeployResources() ([]interface{}, error) {
	return []interface{}{}, nil
}

func (k *cronSource) GenerateWireConnectionResources(previous Component, next Component) ([]interface{}, error) {
	// Next Resource is a channel by expansion rules
	s := map[string]interface{}{
		"apiVersion": k.K8sApiGroup(),
		"kind":       k.K8sKind(),
		"metadata": map[string]interface{}{
			"name":      fmt.Sprintf("%s-%s", k.name, next.K8sName()),
			"namespace": config.Namespace,
		},
		"spec": map[string]interface{}{
			"schedule": k.options["schedule"],
			"data":     k.options["data"],
			"sink":     generateRef(next),
		},
	}
	return []interface{}{s}, nil
}

func (k *cronSource) String() string {
	return fmt.Sprintf("Cron source '%s'", k.name)
}
