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

type task struct {
	config *rest.Config
	gvk    *schema.GroupVersionKind
}

func NewTask(config *rest.Config) *task {
	gvk := &schema.GroupVersionKind{
		Group:   "tekton.dev",
		Kind:    "Task",
		Version: "v1beta1",
	}
	return &task{config: config, gvk: gvk}
}

func (ctl *task) Create(ctx *gin.Context) {
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

func (ctl *task) Get(ctx *gin.Context) {
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

func (ctl *task) ListTask(ctx *gin.Context) {
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

func (ctl *task) Delete(ctx *gin.Context) {
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

func (ctl *task) Put(ctx *gin.Context) {
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
