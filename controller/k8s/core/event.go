package core

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/controller"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type event struct {
	clusterManager k8s.ClusterManager
	gvk            *schema.GroupVersionKind
}

func NewEvent(clusterManager k8s.ClusterManager) *event {
	gvk := &schema.GroupVersionKind{
		Group:   "events.k8s.io",
		Kind:    "event",
		Version: "v1",
	}
	return &event{clusterManager: clusterManager, gvk: gvk}
}

func (ctl *event) ListEvents(ctx *gin.Context) {
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
