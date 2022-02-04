package pipeline

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/controller"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

type trigger struct {
	config *rest.Config
	gvk    *schema.GroupVersionKind
}

func NewTrigger(config *rest.Config) *trigger {
	gvk := &schema.GroupVersionKind{
		Group:   "tekton.dev",
		Kind:    "Trigger",
		Version: "v1beta1",
	}
	return &trigger{config: config, gvk: gvk}
}

func (ctl *trigger) Create(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	obj := &runtime.Unknown{}
	ctx.BindJSON(obj)

	err := k8s.NewResource(ctl.config, ctl.gvk).
		Namespace(namespace).
		Create(obj)

	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *trigger) Get(ctx *gin.Context) {
	name := ctx.Param("name")
	namespace := ctx.Param("namespace")

	obj := &runtime.Unknown{}
	err := k8s.NewResource(ctl.config, ctl.gvk).
		Namespace(namespace).
		Get(name, obj)

	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *trigger) ListTrigger(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	obj := &runtime.Unknown{}
	err := k8s.NewResource(ctl.config, ctl.gvk).
		Namespace(namespace).
		List(obj, metav1.ListOptions{})

	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *trigger) Delete(ctx *gin.Context) {
	name := ctx.Param("name")
	namespace := ctx.Param("namespace")

	err := k8s.NewResource(ctl.config, ctl.gvk).
		Namespace(namespace).
		Delete(name)

	ctx.JSON(200, &controller.Response{
		Err: err,
	})
	return
}

func (ctl *trigger) Put(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")

	obj := &runtime.Unknown{}
	err := k8s.NewResource(ctl.config, ctl.gvk).
		Namespace(namespace).
		Put(name, obj)

	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}
