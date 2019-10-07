package config

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func CreateK8sClientConfig() (*rest.Config, error) {
	if os.Getenv("KFN_IN_CLUSTER") == "true" {
		return rest.InClusterConfig()
	} else {
		if Kubeconfig != "" {
			return clientcmd.BuildConfigFromFlags("", Kubeconfig)
		} else {
			return clientcmd.BuildConfigFromKubeconfigGetter("", clientcmd.NewDefaultClientConfigLoadingRules().Load)
		}
	}
}
