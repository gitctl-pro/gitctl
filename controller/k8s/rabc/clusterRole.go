package rabc

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/controller"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"net/http"
)

type clusterRole struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewClusterRole(clusterManager k8s.ClusterManager) *clusterRole {
	gvk := &schema.GroupVersionKind{
		Group:   "rbac.authorization.k8s.io",
		Kind:    "clusterRole",
		Version: "v1",
	}
	return &clusterRole{clusterManager: clusterManager, gvk: gvk}
}

func (ctl *clusterRole) ListClusterRole(ctx *gin.Context) {
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

func (ctl *clusterRole) Get(ctx *gin.Context) {
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

func (ctl *clusterRole) Put(ctx *gin.Context) {
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

func (ctl *clusterRole) Delete(ctx *gin.Context) {
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

func (ctl *clusterRole) Create(ctx *gin.Context) {
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
