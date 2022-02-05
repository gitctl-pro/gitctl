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

type service struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewService(clusterManager k8s.ClusterManager) *service {
	gvk := &schema.GroupVersionKind{
		Kind:    "service",
		Version: "v1",
	}
	return &service{clusterManager: clusterManager, gvk: gvk}
}

func (ctl *service) ListService(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	namespace := ctx.Param("namespace")

	cfg, _ := ctl.clusterManager.Get(cluster)
	obj := &runtime.Unknown{}
	err := k8s.NewResource(cfg, ctl.gvk).
		Namespace(namespace).
		List(obj, metav1.ListOptions{})

	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *service) Get(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")

	cfg, _ := ctl.clusterManager.Get(cluster)
	obj := &runtime.Unknown{}
	err := k8s.NewResource(cfg, ctl.gvk).
		Namespace(namespace).
		Get(name, obj)

	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *service) Put(ctx *gin.Context) {
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

func (ctl *service) Delete(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")

	cfg, _ := ctl.clusterManager.Get(cluster)
	err := k8s.NewResource(cfg, ctl.gvk).
		Namespace(namespace).
		Delete(name)

	ctx.JSON(200, &controller.Response{
		Err: err,
	})
	return
}

func (ctl *service) Create(ctx *gin.Context) {
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

func (ctl *service) Events(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *service) Pods(ctx *gin.Context) {
	panic("implement me")
}
