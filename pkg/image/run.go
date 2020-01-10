package image

import (
	"github.com/slinkydeveloper/kfn/pkg/config"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	serving_v1 "knative.dev/serving/pkg/apis/serving/v1"
	"knative.dev/serving/pkg/client/clientset/versioned"
)

func (image FunctionImage) RunImage(client *versioned.Clientset, serviceName string) error {
	service := image.ConstructService(serviceName)

	_, err := client.ServingV1().Services(config.Namespace).Create(&service)

	return err
}

// Create service struct from provided options
func (image FunctionImage) ConstructService(name string) serving_v1.Service {
	service := serving_v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: config.Namespace,
		},
	}

	service.Spec.Template = serving_v1.RevisionTemplateSpec{
		Spec: serving_v1.RevisionSpec{},
	}
	service.Spec.Template.Spec.Containers = []corev1.Container{{
		Image: image.FullNameForK8s(),
	}}

	return service
}

func (image FunctionImage) ConstructServiceYaml(name string) map[string]interface{} {
	return map[string]interface{}{
		"apiVersion": "serving.knative.dev/v1alpha1",
		"kind":       "Service",
		"metadata": map[string]string{
			"name":      name,
			"namespace": config.Namespace,
		},
		"spec": map[string]interface{}{
			"template": map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []interface{}{map[string]interface{}{
						"image": image.FullNameForK8s(),
					}},
				},
			},
		},
	}
}
