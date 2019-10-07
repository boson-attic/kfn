package config

import (
	"context"
	"fmt"
	"github.com/containers/buildah/pkg/parse"
	"github.com/containers/image/docker"
	"github.com/containers/image/types"
	routev1typedclient "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	userv1typedclient "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
	"github.com/spf13/viper"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
	containers_config "github.com/containers/image/pkg/docker/config"
)

// To infer the image registry we:
// 1. Try to read image registry from config/env/flag and check if there are credentials
// 2. Try to infer ocp image registry. If there is one available, try to infer credentials from oc whoami
func inferImageRegistry(systemContext *types.SystemContext) (string, string, string, error) {
	return tryDifferentRegistryConfigs(systemContext, inferImageRegistryFromEnv, inferImageRegistryFromOCPImageRegistry)
}

func inferImageRegistryFromEnv(systemContext *types.SystemContext) (string, string, string, error) {
	if viper.IsSet(REGISTRY) && viper.GetString(REGISTRY) != "" {
		registry := viper.GetString(REGISTRY)
		username := viper.GetString(REGISTRY_USERNAME)
		password := viper.GetString(REGISTRY_PASSWORD)
		if err := validCredentials(*systemContext, registry, username, password); err == nil {
			return registry, username, password, nil
		} else {
			return "", "", "", err
		}
	} else {
		return "", "", "", fmt.Errorf("Cannot find flag/env/config entry %s", REGISTRY)
	}
}

func inferImageRegistryFromOCPImageRegistry(systemContext *types.SystemContext) (string, string, string, error) {
	// Try to infer the address
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

	registry := route.Spec.Host

	// Check if user provided credentials
	if viper.IsSet(REGISTRY_USERNAME) {
		username := viper.GetString(REGISTRY_USERNAME)
		password := viper.GetString(REGISTRY_PASSWORD)
		if err := validCredentials(*systemContext, registry, username, password); err == nil {
			return registry, username, password, nil
		} else {
			return "", "", "", err
		}
	}

	// Nope, try to infer credentials
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
		return "", "", "", fmt.Errorf("Trying to use image registry %s but i cannot find the credentials", registry)
	}

	if err := validCredentials(*systemContext, registry, registryUsername, registryPassword); err == nil {
		return registry, registryUsername, registryPassword, nil
	} else {
		return "", "", "", err
	}
}

func tryDifferentRegistryConfigs(systemContext *types.SystemContext, fs ...func(*types.SystemContext) (string, string, string, error)) (string, string, string, error) {
	errors := make([]error, 0)
	for _, f := range fs {
		if name, user, pass, err := f(systemContext); err != nil {
			errors = append(errors, err)
		} else {
			return name, user, pass, nil
		}
	}
	return "", "", "", fmt.Errorf("Cannot infer the image registry: %+v", errors)
}

func validCredentials(systemContext types.SystemContext, registry string, username string, password string) error {
	server := parse.RegistryFromFullName(parse.ScrubServer(registry))
	if username == "" {
		username, password, _ = containers_config.GetAuthentication(&systemContext, server)
	}
	if username != "" {
		systemContext.DockerAuthConfig = &types.DockerAuthConfig{
			Username: username,
			Password: password,
		}
	}
	return docker.CheckAuth(context.TODO(), &systemContext, username, password, server)
}