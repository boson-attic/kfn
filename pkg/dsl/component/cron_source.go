package component

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg/config"
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

func (k cronSource) K8sName() string {
	return k.name
}

func (k *cronSource) Validate() error {
	return nil
}

func (k cronSource) Expand(component Component) Component {
	switch component.(type) {
	case *Function:
		return defaultExpansionChannelFactory("", nil)
	case *knativeService:
		return defaultExpansionChannelFactory("", nil)
	}
	return nil
}

func (k cronSource) CanConnectTo(component Component) bool {
	switch component.(type) {
	case *Function:
		return true
	case *kafkaChannel:
		return true
	case *knativeService:
		return true
	}
	return false
}

func (k cronSource) IsValidWireStart() bool {
	return true
}

func (k cronSource) GenerateDeployResources() ([]interface{}, error) {
	return []interface{}{}, nil
}

func (k cronSource) GenerateWireConnectionResources(previous Component, next Component) ([]interface{}, error) {
	kch := next.(*kafkaChannel)
	s := map[string]interface{}{
		"apiVersion": "sources.eventing.knative.dev/v1alpha1",
		"kind":       "CronJobSource",
		"metadata": map[string]interface{}{
			"name":      fmt.Sprintf("%s-%s", k.name, kch.name),
			"namespace": config.Namespace,
		},
		"spec": map[string]interface{}{
			"schedule": k.options["schedule"],
			"data":     k.options["data"],
			"sink": map[string]interface{}{
				"apiVersion": "messaging.knative.dev/v1alpha1",
				"kind":       "KafkaChannel",
				"name":       k.name,
			},
		},
	}
	return []interface{}{s}, nil
}

func (k cronSource) String() string {
	return fmt.Sprintf("Cron source '%s'", k.name)
}
