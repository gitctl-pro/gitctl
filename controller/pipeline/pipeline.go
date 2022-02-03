package pipeline

import (
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

type pipeline struct {
	config *rest.Config
	gvk    *schema.GroupVersionKind
}

func NewPipeline(config *rest.Config) *pipeline {
	gvk := &schema.GroupVersionKind{
		Group:   "tekton.dev",
		Kind:    "Pipeline",
		Version: "v1beta1",
	}
	return &pipeline{config: config, gvk: gvk}
}

func (p *pipeline) Get(ctx *gin.Context) {
	panic("implement me")
}

func (p *pipeline) Create(ctx *gin.Context) {
	panic("implement me")
}

func (p *pipeline) Put(ctx *gin.Context) {
	panic("implement me")
}

func (p *pipeline) Delete(ctx *gin.Context) {
	panic("implement me")
}

func (p *pipeline) ListPipeline(ctx *gin.Context) {
	panic("implement me")
}
