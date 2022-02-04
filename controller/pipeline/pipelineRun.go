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

type pipelineRun struct {
	config *rest.Config
	gvk    *schema.GroupVersionKind
}

func NewPipelineRun(config *rest.Config) *pipelineRun {
	gvk := &schema.GroupVersionKind{
		Group:   "tekton.dev",
		Kind:    "PipelineRun",
		Version: "v1beta1",
	}
	return &pipelineRun{config: config, gvk: gvk}
}

func (ctl *pipelineRun) Create(ctx *gin.Context) {
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

func (ctl *pipelineRun) Get(ctx *gin.Context) {
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

func (ctl *pipelineRun) ListPipelineRun(ctx *gin.Context) {
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

func (ctl *pipelineRun) Delete(ctx *gin.Context) {
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

func (ctl *pipelineRun) Put(ctx *gin.Context) {
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
