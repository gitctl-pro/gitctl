package clusterapi

import (
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type machineDeployment struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewMachineDeployment(clusterManager k8s.ClusterManager) *machineDeployment {
	gvk := &schema.GroupVersionKind{
		Group:   "clusters.x-k8s.io",
		Kind:    "MachineDeployment",
		Version: "v1beta1",
	}
	return &machineDeployment{clusterManager: clusterManager, gvk: gvk}
}
