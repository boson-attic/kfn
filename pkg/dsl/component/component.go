package component

var anonymousCounter uint

type Component interface {
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
	}
	return nil
}
