package k8s

import (
	"github.com/gin-gonic/gin"
)

type K8s struct {
	Deployment     DeploymentInterface
	Node           NodeInterface
	Namespace      NamespaceInterface
	ReplicaSet     ReplicaSetInterface
	Pod            PodInterface
	DaemonSet      DaemonsetInterface
	ConfigMap      ConfigMapInterface
	Service        ServiceInterface
	Ingress        IngressInterface
	Job            JobInterface
	CronJob        CronJobInterface
	Secret         SecretInterface
	HPA            HPAInterface
	PV             PVInterface
	PVC            PVCInterface
	Event          EventInterface
	Apply          ApplyInterface
	Scale          ScaleInterface
	ServiceAccount ServiceAccountInterface
	ClusterRole    ClusterRoleInterface
}

func NewController() *K8s {
	return &K8s{
		Deployment:     NewDeployment(),
		Namespace:      NewNamespace(),
		Node:           NewNode(),
		Pod:            NewPod(),
		ReplicaSet:     NewReplicaSet(),
		ConfigMap:      NewConfigMap(),
		Service:        NewService(),
		Ingress:        NewIngress(),
		HPA:            NewHPA(),
		Job:            NewJob(),
		CronJob:        NewCronJob(),
		Secret:         NewSecret(),
		PVC:            NewPVC(),
		PV:             NewPV(),
		Event:          NewEvent(),
		Scale:          NewScale(),
		ServiceAccount: NewServiceAccount(),
		ClusterRole:    NewClusterRole(),
	}
}

type DeploymentInterface interface {
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	List(ctx *gin.Context)
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
	List(ctx *gin.Context)
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
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
	Events(ctx *gin.Context)
	Quota(ctx *gin.Context)
	LimitRange(ctx *gin.Context)
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

type ConfigMapInterface interface {
	ListConfigMap(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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
	ListEvent(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Events(ctx *gin.Context)
}

type JobInterface interface {
	ListJob(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Events(ctx *gin.Context)
}

type CronJobInterface interface {
	ListCronJob(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
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
	UpdateHPA(ctx *gin.Context)
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

type ApplyInterface interface {
	Apply(ctx *gin.Context)
}

type ScaleInterface interface {
	ScaleResource(ctx *gin.Context)
}

type ServiceAccountInterface interface {
	ListServiceAccount(ctx *gin.Context)
	Get(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ClusterRoleInterface interface {
	ListClusterRole(ctx *gin.Context)
	Get(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
