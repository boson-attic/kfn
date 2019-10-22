package component

import (
	"fmt"
	"github.com/slinkydeveloper/kfn/pkg/config"
)

func generateChannelToChannelSub(previous Component, actual Component, next Component) map[string]interface{} {
	specMap := map[string]interface{}{
		"channel": generateRef(previous),
		"subscriber": map[string]interface{}{
			"ref": generateRef(actual),
		},
	}

	if next != nil {
		specMap["reply"] = map[string]interface{}{
			"channel": generateRef(next),
		}
	}

	var subName string
	if next != nil {
		subName = fmt.Sprintf("%s-%s-%s", previous.K8sName(), actual.K8sName(), next.K8sName())
	} else {
		subName = fmt.Sprintf("%s-%s", previous.K8sName(), actual.K8sName())
	}

	return map[string]interface{}{
		"apiVersion": MESSAGING_V1ALPHA1_API_GROUP,
		"kind":       "Subscription",
		"metadata": map[string]interface{}{
			"name":      subName,
			"namespace": config.Namespace,
		},
		"spec": specMap,
	}
}
