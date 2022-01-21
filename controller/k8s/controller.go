package k8s

import (
	"github.com/gin-gonic/gin"
)

type K8s struct {
	Application ApplicationInterface
	Deployment  DeploymentInterface
	Cluster     ClusterInterface
	Node        NodeInterface
	Namespace   NamespaceInterface
	ReplicaSet  ReplicaSetInterface
	Pod         PodInterface
	Daemonset   DaemonsetInterface
	Configmap   ConfigmapInterface
	Service     ServiceInterface
	Ingress     IngressInterface
	Job         JobInterface
	CronJob     CronJobInterface
	Secret      SecretInterface
	HPA         HPAInterface
	PVC         PVCInterface
	Event       EventInterface
	Apply       ApplyInterface
	Scale       ScaleInterface
}

func NewController() *K8s {
	return &K8s{
		Deployment:  NewDeployment(),
		Cluster:     NewCluster(),
		Application: NewApplication(),
		Namespace:   NewNamespace(),
		Node:        NewNode(),
		ReplicaSet:  NewReplicaSet(),
		Configmap:   NewConfigmap(),
		Service:     NewService(),
		Ingress:     NewIngress(),
		HPA:         NewHPA(),
		Job:         NewJob(),
		CronJob:     NewCronJob(),
	}
}

type DeploymentInterface interface {
	Events(ctx *gin.Context)
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	ReplicaSets(ctx *gin.Context)
	NewReplicaSets(ctx *gin.Context)
	OldReplicaSets(ctx *gin.Context)
	RolloutRestart(ctx *gin.Context)
	RolloutPause(ctx *gin.Context)
	RolloutResume(ctx *gin.Context)
	RolloutRollback(ctx *gin.Context)
}

type PodInterface interface {
	List(ctx *gin.Context)
	Containers(ctx *gin.Context)
	Get(ctx *gin.Context)
	Events(ctx *gin.Context)
	PersistentVolumeClaims(ctx *gin.Context)
	ExecShell(ctx *gin.Context)
	ExecShellInfo(ctx *gin.Context)
	LogDetail(ctx *gin.Context)
}

type ClusterInterface interface {
	ListCluster(ctx *gin.Context)
	GetCluster(ctx *gin.Context)
}

type NamespaceInterface interface {
	ListNamespace(ctx *gin.Context)
	UpdateNamespace(ctx *gin.Context)
	DeleteNamespace(ctx *gin.Context)
	CreateNamespace(ctx *gin.Context)
}

type ApplicationInterface interface {
	ListApplication(ctx *gin.Context)
	GetApplication(ctx *gin.Context)
	UpdateApplication(ctx *gin.Context)
	DeleteApplication(ctx *gin.Context)
	CreateApplication(ctx *gin.Context)
}

type ReplicaSetInterface interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Events(ctx *gin.Context)
	Pods(ctx *gin.Context)
	Service(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type NodeInterface interface {
	ListNode(ctx *gin.Context)
	GetNode(ctx *gin.Context)
}

type ConfigmapInterface interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Events(ctx *gin.Context)
}

type DaemonsetInterface interface {
	ListDaemonset(ctx *gin.Context)
	GetDaemonset(ctx *gin.Context)
}

type ServiceInterface interface {
	ListService(ctx *gin.Context)
	GetService(ctx *gin.Context)
	Events(ctx *gin.Context)
	Pods(ctx *gin.Context)
}

type IngressInterface interface {
	ListIngress(ctx *gin.Context)
	GetIngress(ctx *gin.Context)
	Events(ctx *gin.Context)
}

type JobInterface interface {
	ListJob(ctx *gin.Context)
	GetJob(ctx *gin.Context)
	Events(ctx *gin.Context)
}

type CronJobInterface interface {
	ListCronJob(ctx *gin.Context)
	GetCronJob(ctx *gin.Context)
	Events(ctx *gin.Context)
}

type SecretInterface interface {
	ListSecret(ctx *gin.Context)
	GetSecret(ctx *gin.Context)
	CreateSecret(ctx *gin.Context)
}

type HPAInterface interface {
	GetHPA(ctx *gin.Context)
	ListHPA(ctx *gin.Context)
	CreateHPA(ctx *gin.Context)
	UpdateHPA(ctx *gin.Context)
	DeleteHPA(ctx *gin.Context)
}

type PVCInterface interface {
	ListPVC(ctx *gin.Context)
	GetPVC(ctx *gin.Context)
	CreatePVC(ctx *gin.Context)
}

type PVInterface interface {
	ListSecret(ctx *gin.Context)
	GetSecret(ctx *gin.Context)
	CreateSecret(ctx *gin.Context)
}

type EventInterface interface {
	ListEvents(ctx *gin.Context)
}

type ApplyInterface interface {
	Apply(ctx *gin.Context)
}

type ScaleInterface interface {
	ScaleResource(ctx *gin.Context)
}
