package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/controller"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type deployment struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewDeployment(clusterManager k8s.ClusterManager) *deployment {
	gvk := &schema.GroupVersionKind{
		Group:   "apps",
		Kind:    "deployment",
		Version: "v1",
	}
	return &deployment{clusterManager: clusterManager, gvk: gvk}
}

func (ctl *deployment) Get(ctx *gin.Context) {
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

func (ctl *deployment) Create(ctx *gin.Context) {
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

func (ctl *deployment) Put(ctx *gin.Context) {
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

func (ctl *deployment) Delete(ctx *gin.Context) {
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

func (ctl *deployment) Patch(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) List(ctx *gin.Context) {
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

func (ctl *deployment) ReplicaSets(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) NewReplicaSets(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) OldReplicaSets(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) RolloutRestart(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) RolloutPause(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) RolloutResume(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) RolloutRollback(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) Events(ctx *gin.Context) {
	panic("implement me")
}
