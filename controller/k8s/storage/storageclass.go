package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/controller"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"net/http"
)

type storageclass struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewStorageClass(clusterManager k8s.ClusterManager) *storageclass {
	gvk := &schema.GroupVersionKind{
		Group:   "storage.k8s.io",
		Kind:    "storageClass",
		Version: "v1",
	}
	return &storageclass{clusterManager: clusterManager, gvk: gvk}
}

func (ctl *storageclass) Create(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	namespace := ctx.Param("namespace")
	obj := &runtime.Unknown{}
	ctx.BindJSON(obj)

	cfg, _ := ctl.clusterManager.Get(cluster)
	err := k8s.NewResource(cfg, ctl.gvk).
		Namespace(namespace).
		Create(obj)

	ctx.JSON(http.StatusCreated, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *storageclass) Get(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	name := ctx.Param("name")

	cfg, _ := ctl.clusterManager.Get(cluster)
	obj := &runtime.Unknown{}
	err := k8s.NewResource(cfg, ctl.gvk).
		Get(name, obj)

	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *storageclass) ListStorageClass(ctx *gin.Context) {
	cluster := ctx.Param("cluster")

	cfg, _ := ctl.clusterManager.Get(cluster)
	obj := &runtime.Unknown{}
	err := k8s.NewResource(cfg, ctl.gvk).
		List(obj, metav1.ListOptions{})

	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *storageclass) Delete(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	name := ctx.Param("name")

	cfg, _ := ctl.clusterManager.Get(cluster)
	err := k8s.NewResource(cfg, ctl.gvk).
		Delete(name)

	ctx.JSON(200, &controller.Response{
		Err: err,
	})
	return
}

func (ctl *storageclass) Put(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")

	cfg, _ := ctl.clusterManager.Get(cluster)
	obj := &runtime.Unknown{}
	err := k8s.NewResource(cfg, ctl.gvk).
		Namespace(namespace).
		Put(name, obj)

	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}
