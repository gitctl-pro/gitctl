package pipeline

import "github.com/gin-gonic/gin"

type pipelineRunCtl struct{}

func NewPipelineRun() PipelineRun {
	return &pipelineRunCtl{}
}

func (p *pipelineRunCtl) GetPipelineRun(ctx *gin.Context) {
	panic("implement me")
}

func (p *pipelineRunCtl) CreatePipelineRun(ctx *gin.Context) {
	panic("implement me")
}

func (p *pipelineRunCtl) DeletePipelineRun(ctx *gin.Context) {
	panic("implement me")
}

func (p *pipelineRunCtl) ListPipelineRun(ctx *gin.Context) {
	panic("implement me")
}
