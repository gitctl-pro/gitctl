package pipeline

import "github.com/gin-gonic/gin"

type pipelineCtl struct{}

func NewPipeline() Pipeline {
	return &pipelineCtl{}
}

func (p *pipelineCtl) GetPipeline(ctx *gin.Context) {
	panic("implement me")
}

func (p *pipelineCtl) CreatePipeline(ctx *gin.Context) {
	panic("implement me")
}

func (p *pipelineCtl) UpdatePipeline(ctx *gin.Context) {
	panic("implement me")
}

func (p *pipelineCtl) DeletePipeline(ctx *gin.Context) {
	panic("implement me")
}

func (p *pipelineCtl) ListPipeline(ctx *gin.Context) {
	panic("implement me")
}
