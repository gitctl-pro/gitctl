package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type pod struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewPod(clusterManager k8s.ClusterManager) PodInterface {
	gvk := &schema.GroupVersionKind{
		Kind:    "pod",
		Version: "v1",
	}
	return &pod{clusterManager: clusterManager, gvk: gvk}
}

func (ctl *pod) ListPod(ctx *gin.Context) {
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

func (ctl *pod) Get(ctx *gin.Context) {
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

func (ctl *pod) Put(ctx *gin.Context) {
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

func (ctl *pod) Delete(ctx *gin.Context) {
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

func (ctl *pod) Create(ctx *gin.Context) {
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

func (ctl *pod) Containers(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pod) Events(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pod) PersistentVolumeClaims(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pod) ExecShell(ctx *gin.Context) {
	panic("implement me")
}
func (ctl *pod) ExecShellInfo(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pod) LogDetail(ctx *gin.Context) {
	panic("implement me")
}
