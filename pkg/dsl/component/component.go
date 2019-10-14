package component

var anonymousCounter uint

type Component interface {
	K8sName() string
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
