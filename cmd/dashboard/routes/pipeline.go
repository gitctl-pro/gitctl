package routes

import (
	"github.com/gitctl-pro/gitctl/controller/pipeline"
)

func (r *RouteManager) addPipelineRoutes(path string) {
	rg := r.gin.Group(path)
	pipeline := pipeline.NewController(nil)
	rg.Use()
	{
		rg.GET("/history", pipeline.History.ListHistories)
		rg.GET("/histroy/:name", pipeline.History.Get)

		rg.GET("/trigger", pipeline.Trigger.ListTrigger)
		rg.GET("/trigger/:name", pipeline.Trigger.Get)
		rg.POST("/trigger/:name", pipeline.Trigger.Create)
		rg.PUT("/trigger/:name", pipeline.Trigger.Put)
		rg.DELETE("/trigger/:name", pipeline.Trigger.Delete)

		rg.GET("/pipeline", pipeline.Trigger.ListTrigger)
		rg.GET("/pipeline/:name", pipeline.Trigger.Get)
		rg.POST("/pipeline/:name", pipeline.Trigger.Create)
		rg.PUT("/pipeline/:name", pipeline.Trigger.Put)
		rg.DELETE("/pipeline/:name", pipeline.Pipeline.Delete)

		rg.GET("/pipelineRun/:name", pipeline.PipelineRun.Get)
		rg.POST("/pipelineRun/:name", pipeline.PipelineRun.Create)
		rg.DELETE("/pipelineRun/:name", pipeline.PipelineRun.Delete)

		rg.GET("/taskRun", pipeline.TaskRun.ListTaskRun)
		rg.GET("/taskRun/:name", pipeline.TaskRun.Get)
		rg.POST("/taskRun/:name", pipeline.TaskRun.Create)
		rg.GET("/taskRun/:name/log", pipeline.TaskRun.Log)

		rg.GET("/task", pipeline.Task.ListTask)
		rg.GET("/task/:name", pipeline.Task.Get)
		rg.POST("/task/:name", pipeline.Task.Create)
		rg.PUT("/task/:name", pipeline.Task.Put)
		rg.DELETE("/task/:name", pipeline.Task.Delete)
	}
}
