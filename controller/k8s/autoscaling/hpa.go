package autoscaling

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/controller"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type hpa struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewHPA(clusterManager k8s.ClusterManager) *hpa {
	gvk := &schema.GroupVersionKind{
		Group:   "autoscaling",
		Kind:    "horizontalpodautoscaler",
		Version: "v2",
	}
	return &hpa{clusterManager: clusterManager, gvk: gvk}
}

func (ctl *hpa) Get(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	name := ctx.Param("name")
	namespace := ctx.Param("namespace")

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

func (ctl *hpa) ListHPA(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	name := ctx.Param("name")
	namespace := ctx.Param("namespace")

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

func (ctl *hpa) Put(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	name := ctx.Param("name")
	namespace := ctx.Param("namespace")

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

func (ctl *hpa) Delete(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	name := ctx.Param("name")
	namespace := ctx.Param("namespace")

	cfg, _ := ctl.clusterManager.Get(cluster)
	err := k8s.NewResource(cfg, ctl.gvk).
		Namespace(namespace).
		Delete(name)

	ctx.JSON(200, &controller.Response{
		Err: err,
	})
	return
}

func (ctl *hpa) Create(ctx *gin.Context) {
	cluster := ctx.Param("cluster")
	namespace := ctx.Param("namespace")

	obj := &runtime.Unknown{}
	ctx.BindJSON(obj)

	cfg, _ := ctl.clusterManager.Get(cluster)
	err := k8s.NewResource(cfg, ctl.gvk).
		Namespace(namespace).
		Create(obj)

	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}
