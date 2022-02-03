package pipeline

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/rest"
)

type PipelineController struct {
	History     History
	Hub         Hub
	Pipeline    Pipeline
	PipelineRun PipelineRun
	Task        Task
	TaskRun     TaskRun
	Trigger     Trigger
}

func NewController(config *rest.Config) *PipelineController {
	return &PipelineController{
		History:     NewHistory(),
		Hub:         NewHub(),
		Pipeline:    NewPipeline(config),
		PipelineRun: NewPipelineRun(config),
		Task:        NewTask(config),
		TaskRun:     NewTaskRun(config),
		Trigger:     NewTrigger(config),
	}
}

type History interface {
	ListHistories(ctx *gin.Context)
	Get(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type Hub interface {
	ListHubs(ctx *gin.Context)
}

type Pipeline interface {
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	ListPipeline(ctx *gin.Context)
}

type PipelineRun interface {
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	ListPipelineRun(ctx *gin.Context)
}

type Task interface {
	ListTask(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type TaskRun interface {
	ListTaskRun(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Log(ctx *gin.Context)
}

type Trigger interface {
	ListTrigger(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
