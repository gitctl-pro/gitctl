package k8s

import "github.com/gin-gonic/gin"

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
	GetLogs(ctx *gin.Context)
	Eviction(ctx *gin.Context)
	Get(ctx *gin.Context)
	Events(ctx *gin.Context)
	PersistentVolumeClaims(ctx *gin.Context)
	ExecShell(ctx *gin.Context)
	ExecShellInfo(ctx *gin.Context)
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

type StatefulSetInterface interface {
	ListStatefulSet(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Events(ctx *gin.Context)
	Pods(ctx *gin.Context)
	Service(ctx *gin.Context)
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
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type HPAInterface interface {
	ListHPA(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type PVCInterface interface {
	ListPVC(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type PVInterface interface {
	ListPV(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
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

type RoleInterface interface {
	ListRole(ctx *gin.Context)
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

type StorageClassInterface interface {
	ListStorageClass(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type NetworkPolicyInterface interface {
	ListNetworkPolicy(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type IngressClassInterface interface {
	ListIngressClass(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}
