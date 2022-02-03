package clusterapi

import (
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type machineHealthCheck struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewMachineHealthCheck(clusterManager k8s.ClusterManager) *machineHealthCheck {
	gvk := &schema.GroupVersionKind{
		Group:   "clusters.x-k8s.io",
		Kind:    "MachineHealthCheck",
		Version: "v1beta1",
	}
	return &machineHealthCheck{clusterManager: clusterManager, gvk: gvk}
}
