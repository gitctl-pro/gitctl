package routes

import (
	"github.com/gitctl-pro/gitctl/controller/pipeline"
)

func (r *RouteManager) addPipelineRoutes(path string) {
	rg := r.gin.Group(path)
	pipeline := pipeline.NewController()
	rg.Use()
	{
		rg.GET("/histories", pipeline.History.ListHistories)
		rg.GET("/histroy/:name", pipeline.History.GetHistory)

		rg.GET("/triggers", pipeline.Trigger.ListTrigger)
		rg.GET("/trigger/:name", pipeline.Trigger.GetTrigger)
		rg.POST("/trigger/:name", pipeline.Trigger.CreateTrigger)
		rg.POST("/trigger/:name", pipeline.Trigger.UpdateTrigger)
		rg.DELETE("/trigger/:name", pipeline.Trigger.DeleteTrigger)

		rg.GET("/pipelines", pipeline.Trigger.ListTrigger)
		rg.GET("/pipeline/:name", pipeline.Trigger.GetTrigger)
		rg.POST("/pipeline/:name", pipeline.Trigger.CreateTrigger)
		rg.POST("/pipeline/:name", pipeline.Trigger.UpdateTrigger)
		rg.DELETE("/pipeline/:name", pipeline.Pipeline.DeletePipeline)

		rg.GET("/pipelineRun/:name", pipeline.PipelineRun.GetPipelineRun)
		rg.POST("/pipelineRun/:name", pipeline.PipelineRun.CreatePipelineRun)
		rg.DELETE("/pipelineRun/:name", pipeline.PipelineRun.DeletePipelineRun)

		rg.GET("/taskRuns", pipeline.TaskRun.ListTaskRun)
		rg.GET("/taskRun/:name", pipeline.TaskRun.GetTaskRun)
		rg.POST("/taskRun/:name", pipeline.TaskRun.CreateTaskRun)
		rg.GET("/taskRun/:name/log", pipeline.TaskRun.Log)

		rg.GET("/tasks", pipeline.Task.ListTask)
		rg.GET("/task/:name", pipeline.Task.GetTask)
		rg.POST("/task/:name", pipeline.Task.CreateTask)
		rg.POST("/task/:name", pipeline.Task.UpdateTask)
		rg.DELETE("/task/:name", pipeline.Task.DeleteTask)
	}
}
