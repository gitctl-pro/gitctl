package core

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/controller"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"net/http"
)

type namespace struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewNamespace(clusterManager k8s.ClusterManager) *namespace {
	gvk := &schema.GroupVersionKind{
		Kind:    "namespace",
		Version: "v1",
	}
	return &namespace{clusterManager: clusterManager, gvk: gvk}
}

func (ctl *namespace) Get(ctx *gin.Context) {
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

func (ctl *namespace) List(ctx *gin.Context) {
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

func (ctl *namespace) Put(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	name := ctx.Param("name")

	cfg, _ := ctl.clusterManager.Get(cluster)
	obj := &runtime.Unknown{}
	err := k8s.NewResource(cfg, ctl.gvk).
		Put(name, obj)

	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *namespace) Delete(ctx *gin.Context) {
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

func (ctl *namespace) Create(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	obj := &runtime.Unknown{}
	ctx.BindJSON(obj)

	cfg, _ := ctl.clusterManager.Get(cluster)
	err := k8s.NewResource(cfg, ctl.gvk).
		Create(obj)

	ctx.JSON(http.StatusCreated, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *namespace) Events(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *namespace) Quota(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *namespace) LimitRange(ctx *gin.Context) {
	panic("implement me")
}
