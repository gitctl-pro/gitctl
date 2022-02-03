package clusterapi

import (
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type cluster struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewCluster(clusterManager k8s.ClusterManager) *cluster {
	gvk := &schema.GroupVersionKind{
		Group:   "clusters.x-k8s.io",
		Kind:    "Cluster",
		Version: "v1beta1",
	}
	return &cluster{clusterManager: clusterManager, gvk: gvk}
}
