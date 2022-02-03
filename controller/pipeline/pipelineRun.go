package pipeline

import (
	"github.com/gin-gonic/gin"
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

func (p *pipelineRun) Get(ctx *gin.Context) {
	panic("implement me")
}

func (p *pipelineRun) Create(ctx *gin.Context) {
	panic("implement me")
}

func (p *pipelineRun) Delete(ctx *gin.Context) {
	panic("implement me")
}

func (p *pipelineRun) ListPipelineRun(ctx *gin.Context) {
	panic("implement me")
}
