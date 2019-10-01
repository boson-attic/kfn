package image

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	serving_v1alpha1_api "knative.dev/serving/pkg/apis/serving/v1alpha1"
	servingv1alpha1 "knative.dev/serving/pkg/client/clientset/versioned/typed/serving/v1alpha1"
)

func (image FunctionImage) RunImage(client servingv1alpha1.ServingV1alpha1Interface, serviceName string, namespace string) error {
	service := image.constructService(serviceName, namespace)

	_, err := client.Services(namespace).Create(&service)

	return err
}

// Create service struct from provided options
func (image FunctionImage) constructService(name string, namespace string) serving_v1alpha1_api.Service {
	service := serving_v1alpha1_api.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}

	service.Spec.Template = &serving_v1alpha1_api.RevisionTemplateSpec{
		Spec: serving_v1alpha1_api.RevisionSpec{},
	}
	service.Spec.Template.Spec.Containers = []corev1.Container{{
		Image: image.FullNameForK8s(),
	}}

	return service
}
