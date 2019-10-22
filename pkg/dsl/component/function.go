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

func (f *Function) K8sApiGroup() string {
	return SERVING_V1ALPHA1_API_GROUP
}

func (f *Function) K8sKind() string {
	return "Service"
}

func (f *Function) ComponentType() ComponentType {
	return Service
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
	if component.ComponentType() == Service {
		return defaultExpansionChannelFactory("", nil)
	}
	return nil
}

func (f *Function) CanConnectTo(component Component) bool {
	return util.AnyOf(component.ComponentType(), Channel, Service)
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
	// Previous can't be null because function is not a valid wire start
	if previous.ComponentType() == Channel && (next == nil || next.ComponentType() == Channel) {
		return []interface{}{generateChannelToChannelSub(previous, f, next)}, nil
	}
	return []interface{}{}, nil
}

func (f *Function) String() string {
	return fmt.Sprintf("Function '%s'", f.functionLocation)
}
