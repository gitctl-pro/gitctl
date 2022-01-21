package pipeline

import "github.com/gin-gonic/gin"

type PipelineController struct {
	History     History
	Hub         Hub
	Pipeline    Pipeline
	PipelineRun PipelineRun
	Task        Task
	TaskRun     TaskRun
	Trigger     Trigger
}

func NewController() *PipelineController {
	return &PipelineController{
		History:     NewHistory(),
		Hub:         NewHub(),
		Pipeline:    NewPipeline(),
		PipelineRun: NewPipelineRun(),
		Task:        NewTask(),
		TaskRun:     NewTaskRun(),
		Trigger:     NewTrigger(),
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

type Pipeline interface {
	GetPipeline(ctx *gin.Context)
	CreatePipeline(ctx *gin.Context)
	UpdatePipeline(ctx *gin.Context)
	DeletePipeline(ctx *gin.Context)
	ListPipeline(ctx *gin.Context)
}

type PipelineRun interface {
	GetPipelineRun(ctx *gin.Context)
	CreatePipelineRun(ctx *gin.Context)
	DeletePipelineRun(ctx *gin.Context)
	ListPipelineRun(ctx *gin.Context)
}

type Task interface {
	GetTask(ctx *gin.Context)
	CreateTask(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
	ListTask(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
}

type TaskRun interface {
	ListTaskRun(ctx *gin.Context)
	GetTaskRun(ctx *gin.Context)
	CreateTaskRun(ctx *gin.Context)
	Log(ctx *gin.Context)
}

type Trigger interface {
	GetTrigger(ctx *gin.Context)
	CreateTrigger(ctx *gin.Context)
	UpdateTrigger(ctx *gin.Context)
	ListTrigger(ctx *gin.Context)
	DeleteTrigger(ctx *gin.Context)
}
