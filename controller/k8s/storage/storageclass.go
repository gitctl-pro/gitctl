package rabc

import (
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type storageclass struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewStorageclass(clusterManager k8s.ClusterManager) *storageclass {
	gvk := &schema.GroupVersionKind{
		Group:   "storage.k8s.io",
		Kind:    "storageclass",
		Version: "v1",
	}
	return &storageclass{clusterManager: clusterManager, gvk: gvk}
}
