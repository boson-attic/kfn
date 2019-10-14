package component

import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/slinkydeveloper/kfn/pkg"
	"github.com/slinkydeveloper/kfn/pkg/config"
	"github.com/slinkydeveloper/kfn/pkg/image"
	"github.com/slinkydeveloper/kfn/pkg/languages"
	"github.com/slinkydeveloper/kfn/pkg/util"
	"path"
	"path/filepath"
	"strings"
)

type Function struct {
	functionLocation string
	options          map[string]string
	imageName        string
	serviceName      string
	imageTag         string
	functionImage    *image.FunctionImage
	language         languages.Language
}

func NewFunction(functionLocation string, options map[string]string) Component {
	return &Function{functionLocation: functionLocation, options: options}
}

func (f *Function) K8sName() string {
	return f.serviceName
}

func (f *Function) Validate() error {
	var err error
	f.functionLocation, err = filepath.Abs(f.functionLocation)
	if err != nil {
		return err
	}

	if !util.FsExist(f.functionLocation) {
		return fmt.Errorf("Cannot retrieve function %s", f.functionLocation)
	}

	f.language = languages.GetLanguage(path.Ext(f.functionLocation))
	if f.language == languages.Unknown {
		return fmt.Errorf("unknown language for function %s", f.functionLocation)
	}

	f.imageName = f.options["image_name"]
	f.serviceName = f.options["service_name"]
	f.imageTag = f.options["image_tag"]

	if len(f.imageName) == 0 {
		base := path.Base(f.functionLocation)
		f.imageName = strings.TrimSuffix(base, path.Ext(base))
	}

	if len(f.serviceName) == 0 {
		f.serviceName = f.imageName
	}

	return nil
}

func (f *Function) Expand(component Component) Component {
	switch component.(type) {
	case *Function:
		return defaultExpansionChannelFactory("", nil)
	case *knativeService:
		return defaultExpansionChannelFactory("", nil)
	}
	return nil
}

func (f *Function) CanConnectTo(component Component) bool {
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

func (f *Function) Build() error {
	log.Infof("Starting building %v", f)

	functionImage, err := pkg.Build(f.functionLocation, f.language, f.imageName, f.imageTag, config.BuildSystemContext)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error while trying to build %v", f))
	}

	log.Infof("Completed building %v. Pushed to %s", f, functionImage.FullNameForK8s())

	f.functionImage = &functionImage
	return nil
}

func (f *Function) IsValidWireStart() bool {
	return false
}

func (f *Function) GenerateDeployResources() ([]interface{}, error) {
	return []interface{}{f.functionImage.ConstructServiceYaml(f.serviceName)}, nil
}

func (f *Function) GenerateWireConnectionResources(previous Component, next Component) ([]interface{}, error) {
	switch previous.(type) {
	case *inMemoryChannel:
		if next != nil {
			switch next.(type) {
			case *kafkaChannel:
				return []interface{}{f.generateChannelToChannelSub(previous.K8sName(), "InMemoryChannel", next.K8sName(), "KafkaChannel")}, nil
			case *inMemoryChannel:
				return []interface{}{f.generateChannelToChannelSub(previous.K8sName(), "InMemoryChannel", next.K8sName(), "InMemoryChannel")}, nil
			}
		} else {
			return []interface{}{f.generateChannelToChannelSub(previous.K8sName(), "InMemoryChannel", "", "")}, nil
		}
	case *kafkaChannel:
		if next != nil {
			switch next.(type) {
			case *kafkaChannel:
				return []interface{}{f.generateChannelToChannelSub(previous.K8sName(), "KafkaChannel", next.K8sName(), "KafkaChannel")}, nil
			case *inMemoryChannel:
				return []interface{}{f.generateChannelToChannelSub(previous.K8sName(), "KafkaChannel", next.K8sName(), "InMemoryChannel")}, nil
			}
		} else {
			return []interface{}{f.generateChannelToChannelSub(previous.K8sName(), "KafkaChannel", "", "")}, nil
		}
	}
	return []interface{}{}, nil
}

func (f *Function) generateChannelToChannelSub(previousChannelName string, previousChannelType string, nextChannelName string, nextChannelType string) map[string]interface{} {
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

	return map[string]interface{}{
		"apiVersion": "messaging.knative.dev/v1alpha1",
		"kind":       "Subscription",
		"metadata": map[string]interface{}{
			"name":      fmt.Sprintf("%s-%s-%s", previousChannelName, f.serviceName, nextChannelName),
			"namespace": config.Namespace,
		},
		"spec": specMap,
	}
}

func (f *Function) String() string {
	return fmt.Sprintf("Function '%s'", f.functionLocation)
}
