package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type crd struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewCrd(clusterManager k8s.ClusterManager) CrdInterface {
	gvk := &schema.GroupVersionKind{
		Group:   "apiextensions.k8s.io",
		Kind:    "CustomResourceDefinition",
		Version: "v1",
	}
	return &crd{clusterManager: clusterManager, gvk: gvk}
}

func (ctl *crd) ListCrd(ctx *gin.Context) {
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

func (ctl *crd) Get(ctx *gin.Context) {
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

func (ctl *crd) Put(ctx *gin.Context) {
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

func (ctl *crd) Delete(ctx *gin.Context) {
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

func (ctl *crd) Create(ctx *gin.Context) {
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
