package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type replicaset struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewReplicaset(clusterManager k8s.ClusterManager) ReplicasetInterface {
	gvk := &schema.GroupVersionKind{
		Group:   "apps",
		Kind:    "replicaset",
		Version: "v1",
	}
	return &replicaset{clusterManager: clusterManager, gvk: gvk}
}

func (ctl *replicaset) Create(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	namespace := ctx.Param("namespace")
	obj := &runtime.Unknown{}
	ctx.BindJSON(obj)

	cfg, _ := ctl.clusterManager.Get(cluster)
	err := k8s.NewResource(cfg, ctl.gvk).
		Namespace(namespace).
		Create(obj)

	ctx.JSON(200, &response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *replicaset) Get(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")

	cfg, _ := ctl.clusterManager.Get(cluster)
	obj := &runtime.Unknown{}
	err := k8s.NewResource(cfg, ctl.gvk).
		Namespace(namespace).
		Get(name, obj)

	ctx.JSON(200, &response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *replicaset) Put(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")

	cfg, _ := ctl.clusterManager.Get(cluster)
	obj := &runtime.Unknown{}
	err := k8s.NewResource(cfg, ctl.gvk).
		Namespace(namespace).
		Put(name, obj)

	ctx.JSON(200, &response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *replicaset) Delete(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")

	cfg, _ := ctl.clusterManager.Get(cluster)
	err := k8s.NewResource(cfg, ctl.gvk).
		Namespace(namespace).
		Delete(name)

	ctx.JSON(200, &response{
		Err: err,
	})
	return
}

func (ctl *replicaset) List(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	namespace := ctx.Param("namespace")

	cfg, _ := ctl.clusterManager.Get(cluster)
	obj := &runtime.Unknown{}
	err := k8s.NewResource(cfg, ctl.gvk).
		Namespace(namespace).
		List(obj, metav1.ListOptions{})

	ctx.JSON(200, &response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *replicaset) Events(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *replicaset) Pods(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *replicaset) Service(ctx *gin.Context) {
	panic("implement me")
}
