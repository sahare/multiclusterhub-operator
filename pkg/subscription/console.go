// Copyright (c) 2020 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package subscription

import (
	"os"

	operatorsv1 "github.com/stolostron/multiclusterhub-operator/api/v1"
	"github.com/stolostron/multiclusterhub-operator/pkg/utils"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Console overrides the console-chart chart
func Console(m *operatorsv1.MultiClusterHub, overrides map[string]string, ingress string) *unstructured.Unstructured {
	sub := &Subscription{
		Name:      "console-chart",
		Namespace: m.Namespace,
		Overrides: map[string]interface{}{
			"pullSecret":   m.Spec.ImagePullSecret,
			"ocpingress":   ingress,
			"cfcRouterUrl": "https://management-ingress:443",
			"hubconfig": map[string]interface{}{
				"name":         m.Name,
				"namespace":    m.Namespace,
				"replicaCount": utils.DefaultReplicaCount(m),
				"nodeSelector": m.Spec.NodeSelector,
				"tolerations":  utils.GetTolerations(m),
				"ocpVersion":   os.Getenv("ACM_HUB_OCP_VERSION"),
			},
			"global": map[string]interface{}{
				"imageOverrides": overrides,
				"pullPolicy":     utils.GetImagePullPolicy(m),
			},
			"clusterImageSets": map[string]interface{}{
				"subscriptionPause": utils.GetDisableClusterImageSets(m),
			},
		},
	}
	setCustomCA(m, sub)

	return newSubscription(m, sub)
}
