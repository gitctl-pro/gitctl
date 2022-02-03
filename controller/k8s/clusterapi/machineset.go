package clusterapi

import (
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type machineSet struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewMachineSet(clusterManager k8s.ClusterManager) *machineSet {
	gvk := &schema.GroupVersionKind{
		Group:   "clusters.x-k8s.io",
		Kind:    "MachineSet",
		Version: "v1beta1",
	}
	return &machineSet{clusterManager: clusterManager, gvk: gvk}
}
