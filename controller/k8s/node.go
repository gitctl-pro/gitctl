package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type node struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewNode(clusterManager k8s.ClusterManager) NodeInterface {
	gvk := &schema.GroupVersionKind{
		Group:   "",
		Kind:    "node",
		Version: "v1",
	}
	return &node{clusterManager: clusterManager, gvk: gvk}
}

func (ctl *node) ListNode(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	name := ctx.Param("name")

	cfg, _ := ctl.clusterManager.Get(cluster)
	obj := &runtime.Unknown{}
	err := k8s.NewResource(cfg, ctl.gvk).
		Put(name, obj)

	ctx.JSON(200, &response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *node) Get(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	name := ctx.Param("name")

	cfg, _ := ctl.clusterManager.Get(cluster)
	obj := &runtime.Unknown{}
	err := k8s.NewResource(cfg, ctl.gvk).
		Get(name, obj)

	ctx.JSON(200, &response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *node) Put(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	name := ctx.Param("name")

	cfg, _ := ctl.clusterManager.Get(cluster)
	obj := &runtime.Unknown{}
	err := k8s.NewResource(cfg, ctl.gvk).
		Put(name, obj)

	ctx.JSON(200, &response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *node) Delete(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	name := ctx.Param("name")

	cfg, _ := ctl.clusterManager.Get(cluster)
	err := k8s.NewResource(cfg, ctl.gvk).
		Delete(name)

	ctx.JSON(200, &response{
		Err: err,
	})
	return
}
