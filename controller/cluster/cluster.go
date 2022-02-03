package cluster

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/apps/apis/core/v1"
	"github.com/gitctl-pro/gitctl/pkg/controller"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

type cluster struct {
	config   *rest.Config
	resource k8s.Resource
}

func (ctl *cluster) List(ctx *gin.Context) {
	list := &v1.ClusterList{}
	err := ctl.resource.List(list, metav1.ListOptions{})
	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: list,
	})
	return
}

func (ctl *cluster) Get(ctx *gin.Context) {
	name := ctx.Param("name")
	obj := &v1.Cluster{}
	err := ctl.resource.Get(name, obj)
	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *cluster) Put(ctx *gin.Context) {
	name := ctx.Param("name")
	obj := &v1.Cluster{}
	ctx.BindJSON(obj)
	err := ctl.resource.Put(name, obj)
	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}

func (ctl *cluster) Delete(ctx *gin.Context) {
	name := ctx.Param("name")
	err := ctl.resource.Delete(name)
	ctx.JSON(200, &controller.Response{
		Err: err,
	})
	return
}

func (ctl *cluster) Create(ctx *gin.Context) {
	obj := &v1.Cluster{}
	ctx.BindJSON(obj)
	err := ctl.resource.Create(obj)
	ctx.JSON(200, &controller.Response{
		Err:  err,
		Data: obj,
	})
	return
}
