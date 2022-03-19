package k8s

import (
	"github.com/gitctl-pro/gitctl/controller/k8s/apps"
	"github.com/gitctl-pro/gitctl/controller/k8s/autoscaling"
	"github.com/gitctl-pro/gitctl/controller/k8s/batch"
	"github.com/gitctl-pro/gitctl/controller/k8s/core"
	"github.com/gitctl-pro/gitctl/controller/k8s/extension"
	"github.com/gitctl-pro/gitctl/controller/k8s/networking"
	"github.com/gitctl-pro/gitctl/controller/k8s/rabc"
	"github.com/gitctl-pro/gitctl/controller/k8s/storage"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/client-go/rest"
)

type K8sController struct {
	clusterManager k8s.ClusterManager
	Deployment     DeploymentInterface
	StatefulSet    StatefulSetInterface
	DaemonSet      DaemonsetInterface
	ReplicaSet     ReplicasetInterface
	Node           NodeInterface
	Namespace      NamespaceInterface
	Pod            PodInterface
	ConfigMap      ConfigmapInterface
	Service        ServiceInterface
	Ingress        IngressInterface
	IngressClass   IngressClassInterface
	NetworkPolicy  NetworkPolicyInterface
	Job            JobInterface
	CronJob        CronjobInterface
	Secret         SecretInterface
	HPA            HPAInterface
	PV             PVInterface
	PVC            PVCInterface
	Event          EventInterface
	Scale          ScaleInterface
	ServiceAccount ServiceAccountInterface
	ClusterRole    ClusterRoleInterface
	Role           RoleInterface
	Crd            CrdInterface
	StorageClass   StorageClassInterface
}

func NewController(cfg *rest.Config, clusterManager k8s.ClusterManager) *K8sController {
	return &K8sController{
		Deployment:     apps.NewDeployment(clusterManager),
		ReplicaSet:     apps.NewReplicaset(clusterManager),
		DaemonSet:      apps.NewDaemonset(clusterManager),
		StatefulSet:    apps.NewStatefulSet(clusterManager),
		Namespace:      core.NewNamespace(clusterManager),
		Node:           core.NewNode(clusterManager),
		Pod:            core.NewPod(clusterManager),
		ConfigMap:      core.NewConfigmap(clusterManager),
		Service:        core.NewService(clusterManager),
		Ingress:        networking.NewIngress(clusterManager),
		IngressClass:   networking.NewIngressClass(clusterManager),
		NetworkPolicy:  networking.NewNetworkPolicy(clusterManager),
		HPA:            autoscaling.NewHPA(clusterManager),
		Job:            batch.NewJob(clusterManager),
		CronJob:        batch.NewCronjob(clusterManager),
		Secret:         core.NewSecret(clusterManager),
		PVC:            core.NewPVC(clusterManager),
		PV:             core.NewPV(clusterManager),
		Event:          core.NewEvent(clusterManager),
		ServiceAccount: core.NewServiceAccount(clusterManager),
		ClusterRole:    rabc.NewClusterRole(clusterManager),
		Role:           rabc.NewRole(clusterManager),
		Crd:            extension.NewCrd(clusterManager),
		StorageClass:   storage.NewStorageClass(clusterManager),
	}
}
