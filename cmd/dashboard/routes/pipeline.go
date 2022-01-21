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
	}
}
