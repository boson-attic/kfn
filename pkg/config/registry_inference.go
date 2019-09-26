package config

import (
	"fmt"
	routev1typedclient "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	userv1typedclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

// To infer the image registry we:
// 1. Try to read image registry from config/env/flag and check if there are credentials
// 2. Try to infer ocp image registry. If there is one available, try to infer credentials from oc whoami
func inferImageRegistry() (string, string, string, error) {
	return imTryingToDoSomeFPInGoButItLooksAwful(inferImageRegistryFromEnv, inferImageRegistryFromOCPImageRegistry)
}

func inferImageRegistryFromEnv() (string, string, string, error) {
	if viper.IsSet(REGISTRY) && viper.GetString(REGISTRY) != "" {
		return viper.GetString(REGISTRY), viper.GetString(REGISTRY_USERNAME), viper.GetString(REGISTRY_PASSWORD), nil
	} else {
		return "", "", "", fmt.Errorf("Cannot find flag/env/config entry %s", REGISTRY)
	}
}

func inferImageRegistryFromOCPImageRegistry() (string, string, string, error) {
	config, err := CreateK8sClientConfig()
	if err != nil {
		return "", "", "", err
	}
	routeClient, err := routev1typedclient.NewForConfig(config)
	if err != nil {
		return "", "", "", err
	}

	route, err := routeClient.Routes("openshift-image-registry").Get("default-route", v1.GetOptions{})
	if err != nil {
		return "", "", "", err
	}

	registryHost := route.Spec.Host

	userClient, err := userv1typedclient.NewForConfig(config)
	if err != nil {
		return "", "", "", err
	}

	me, err := userClient.Users().Get("~", v1.GetOptions{})
	if err != nil {
		return "", "", "", err
	}
	registryUsername := strings.Trim(me.Name, " ")
	if registryUsername == "kube:admin" {
		registryUsername = "kubeadmin" // Small workaround for this inconsistency
	}

	var registryPassword string
	if config.BearerToken != "" {
		registryPassword = strings.Trim(config.BearerToken, " ")
	} else {
		return "", "", "", fmt.Errorf("Trying to use image registry %s but i cannot find the credentials", registryHost)
	}
	log.Infof("Found image registry '%s' with username %s and pass %s", registryHost, registryUsername, registryPassword)
	return registryHost, registryUsername, registryPassword, nil
}

func imTryingToDoSomeFPInGoButItLooksAwful(fs ...func() (string, string, string, error)) (string, string, string, error) {
	errors := make([]error, 0)
	for _, f := range fs {
		if name, user, pass, err := f(); err != nil {
			errors = append(errors, err)
		} else {
			return name, user, pass, nil
		}
	}
	return "", "", "", fmt.Errorf("Cannot infer the image registry: %+v", errors)
}
