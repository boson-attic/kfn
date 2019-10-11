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
	functionImage    *image.FunctionImage
}

func NewFunction(functionLocation string, options map[string]string) Component {
	return &Function{functionLocation, options, nil}
}

func (f Function) Validate() error {
	if !util.FsExist(f.functionLocation) {
		return fmt.Errorf("Cannot retrieve function %s", f.functionLocation)
	}
	return nil
}

func (f Function) Expand(component Component) Component {
	switch component.(type) {
	case *Function:
		return NewKafkaChannel("", nil)
	}
	return nil
}

func (f Function) CanConnectTo(component Component) bool {
	switch component.(type) {
	case *Function:
		return true
	case *kafkaChannel:
		return true
	}
	return false
}

func (f *Function) Build() error {
	log.Infof("Starting building %v", f)

	functionPath := f.functionLocation
	functionPath, err := filepath.Abs(functionPath)
	if err != nil {
		return err
	}

	language := languages.GetLanguage(path.Ext(functionPath))
	if language == languages.Unknown {
		return fmt.Errorf("unknown language for function %s", functionPath)
	}

	imageName := f.options["image_name"]
	serviceName := f.options["service_name"]
	imageTag := f.options["image_tag"]

	if len(imageName) == 0 {
		base := path.Base(functionPath)
		imageName = strings.TrimSuffix(base, path.Ext(base))
	}

	if len(serviceName) == 0 {
		serviceName = imageName
	}

	f.options["image_name"] = imageName
	f.options["service_name"] = serviceName
	f.options["image_tag"] = imageTag

	functionImage, err := pkg.Build(functionPath, language, imageName, imageTag, config.BuildSystemContext)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error while trying to build %v", f))
	}

	log.Infof("Completed building %v. Pushed to %s", f, functionImage.FullNameForK8s())

	f.functionImage = &functionImage
	return nil
}

func (f Function) IsValidWireStart() bool {
	return false
}

func (f Function) GenerateDeployResources() ([]interface{}, error) {
	return []interface{}{f.functionImage.ConstructServiceYaml(f.options["service_name"])}, nil
}

func (f Function) GenerateWireConnectionResources(previous Component, next Component) ([]interface{}, error) {
	switch previous.(type) {
	case *kafkaChannel:
		if next != nil {
			switch next.(type) {
			case *kafkaChannel:
				return []interface{}{f.generateChannelToChannelSub(previous.(*kafkaChannel), next.(*kafkaChannel))}, nil
			}
		} else {
			return []interface{}{f.generateChannelSub(previous.(*kafkaChannel))}, nil
		}
	}
	return []interface{}{}, nil
}

func (f Function) generateChannelToChannelSub(previousChannel *kafkaChannel, nextChannel *kafkaChannel) map[string]interface{} {
	serviceName := f.options["service_name"]
	return map[string]interface{}{
		"apiVersion": "messaging.knative.dev/v1alpha1",
		"kind":       "Subscription",
		"metadata": map[string]interface{}{
			"name":      fmt.Sprintf("%s-%s-%s", previousChannel.name, serviceName, nextChannel.name),
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
					"name":       serviceName,
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

func (f Function) generateChannelSub(previousChannel *kafkaChannel) map[string]interface{} {
	serviceName := f.options["service_name"]
	return map[string]interface{}{
		"apiVersion": "messaging.knative.dev/v1alpha1",
		"kind":       "Subscription",
		"metadata": map[string]interface{}{
			"name":      fmt.Sprintf("%s-%s", previousChannel.name, serviceName),
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
					"name":       serviceName,
				},
			},
		},
	}
}

func (f Function) String() string {
	return fmt.Sprintf("Function '%s'", f.functionLocation)
}
