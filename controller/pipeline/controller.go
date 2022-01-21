package pipeline

import "github.com/gin-gonic/gin"

type PipelineController struct {
	History History
	Hub     Hub
}

func NewController() *PipelineController {
	return &PipelineController{
		History: NewHistory(),
		Hub:     NewHub(),
	}
}

type History interface {
	GetHistory(ctx *gin.Context)
	ListHistories(ctx *gin.Context)
	DeleteHistory(ctx *gin.Context)
}

type Hub interface {
	ListHubs(ctx *gin.Context)
}
