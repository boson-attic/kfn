package component

var anonymousCounter uint

type ComponentType uint

const (
	Channel ComponentType = iota
	Service
	Source
)

const (
	SOURCES_v1ALPHA1_API_GROUP   = "sources.eventing.knative.dev/v1alpha1"
	MESSAGING_V1ALPHA1_API_GROUP = "messaging.knative.dev/v1alpha1"
	SERVING_V1ALPHA1_API_GROUP   = "serving.knative.dev/v1alpha1"
)

type Component interface {
	K8sName() string
	K8sApiGroup() string
	K8sKind() string
	ComponentType() ComponentType
	Validate() error
	CanConnectTo(component Component) bool
	IsValidWireStart() bool
	Expand(next Component) Component
	GenerateDeployResources() ([]interface{}, error)
	GenerateWireConnectionResources(previous Component, next Component) ([]interface{}, error)
}

func ResolveComponentFactory(t string) func(string, map[string]string) Component {
	switch t {
	case "Function":
		return NewFunction
	case "KafkaChannel":
		return NewKafkaChannel
	case "CronSource":
		return NewCronSource
	case "KService":
		return NewKnativeService
	case "KnativeService":
		return NewKnativeService
	case "InMemoryChannel":
		return NewInMemoryChannel
	}
	return nil
}

var defaultExpansionChannelFactory = NewInMemoryChannel

func generateRef(component Component) map[string]string {
	return map[string]string{
		"apiVersion": component.K8sApiGroup(),
		"kind":       component.K8sKind(),
		"name":       component.K8sName(),
	}
}
