package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/client-go/rest"
)

type K8s struct {
	Cluster        ClusterInterface
	Deployment     DeploymentInterface
	Node           NodeInterface
	Namespace      NamespaceInterface
	ReplicaSet     ReplicasetInterface
	Pod            PodInterface
	DaemonSet      DaemonsetInterface
	ConfigMap      ConfigmapInterface
	Service        ServiceInterface
	Ingress        IngressInterface
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
	Crd            CrdInterface
}

func NewController(cfg *rest.Config, clusterManager k8s.ClusterManager) *K8s {
	return &K8s{
		Cluster:        NewCluster(cfg),
		Deployment:     NewDeployment(clusterManager),
		Namespace:      NewNamespace(clusterManager),
		Node:           NewNode(clusterManager),
		Pod:            NewPod(clusterManager),
		ReplicaSet:     NewReplicaset(clusterManager),
		ConfigMap:      NewConfigmap(clusterManager),
		Service:        NewService(clusterManager),
		Ingress:        NewIngress(clusterManager),
		HPA:            NewHPA(),
		Job:            NewJob(clusterManager),
		CronJob:        NewCronjob(clusterManager),
		Secret:         NewSecret(clusterManager),
		PVC:            NewPVC(),
		PV:             NewPV(),
		Event:          NewEvent(),
		Scale:          NewScale(),
		ServiceAccount: NewServiceAccount(clusterManager),
		ClusterRole:    NewClusterRole(clusterManager),
		Crd:            NewCrd(clusterManager),
	}
}

type response struct {
	Err  error       `json:"err,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type DeploymentInterface interface {
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	List(ctx *gin.Context)
	Patch(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Events(ctx *gin.Context)
	ReplicaSets(ctx *gin.Context)
	NewReplicaSets(ctx *gin.Context)
	OldReplicaSets(ctx *gin.Context)
	RolloutRestart(ctx *gin.Context)
	RolloutPause(ctx *gin.Context)
	RolloutResume(ctx *gin.Context)
	RolloutRollback(ctx *gin.Context)
}

type PodInterface interface {
	ListPod(ctx *gin.Context)
	Containers(ctx *gin.Context)
	Get(ctx *gin.Context)
	Events(ctx *gin.Context)
	PersistentVolumeClaims(ctx *gin.Context)
	ExecShell(ctx *gin.Context)
	ExecShellInfo(ctx *gin.Context)
	LogDetail(ctx *gin.Context)
}

type NamespaceInterface interface {
	List(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
	Events(ctx *gin.Context)
	Quota(ctx *gin.Context)
	LimitRange(ctx *gin.Context)
}

type ReplicasetInterface interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Events(ctx *gin.Context)
	Pods(ctx *gin.Context)
	Service(ctx *gin.Context)
}

type NodeInterface interface {
	ListNode(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ConfigmapInterface interface {
	ListConfigmap(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type DaemonsetInterface interface {
	ListDaemonset(ctx *gin.Context)
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Events(ctx *gin.Context)
	Pods(ctx *gin.Context)
}

type ServiceInterface interface {
	ListService(ctx *gin.Context)
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Events(ctx *gin.Context)
	Pods(ctx *gin.Context)
}

type IngressInterface interface {
	ListIngress(ctx *gin.Context)
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Events(ctx *gin.Context)
}

type JobInterface interface {
	ListJob(ctx *gin.Context)
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Events(ctx *gin.Context)
}

type CronjobInterface interface {
	ListCronjob(ctx *gin.Context)
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Events(ctx *gin.Context)
}

type SecretInterface interface {
	ListSecret(ctx *gin.Context)
	Get(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type HPAInterface interface {
	GetHPA(ctx *gin.Context)
	ListHPA(ctx *gin.Context)
	CreateHPA(ctx *gin.Context)
	PutHPA(ctx *gin.Context)
	DeleteHPA(ctx *gin.Context)
}

type PVCInterface interface {
	ListPVC(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type PVInterface interface {
	ListPV(ctx *gin.Context)
	GetPV(ctx *gin.Context)
	CreatePV(ctx *gin.Context)
	DeletePV(ctx *gin.Context)
}

type EventInterface interface {
	ListEvents(ctx *gin.Context)
}

type ScaleInterface interface {
	ScaleResource(ctx *gin.Context)
}

type ServiceAccountInterface interface {
	ListServiceAccount(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type ClusterRoleInterface interface {
	ListClusterRole(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type ClusterInterface interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type CrdInterface interface {
	ListCrd(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}
