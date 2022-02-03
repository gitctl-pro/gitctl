package clusterapi

import (
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type machine struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewMachine(clusterManager k8s.ClusterManager) *machine {
	gvk := &schema.GroupVersionKind{
		Group:   "clusters.x-k8s.io",
		Kind:    "Machine",
		Version: "v1beta1",
	}
	return &machine{clusterManager: clusterManager, gvk: gvk}
}
