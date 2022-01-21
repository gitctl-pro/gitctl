package routes

import (
	"github.com/gitctl-pro/gitctl/controller/cluster"
)

func (r *RouteManager) addClusterRoutes(path string) {
	rg := r.gin.Group(path)
	cluster := cluster.NewController()
	// route: /api/v1/cluster
	rg = r.gin.Group(path + "/application")
	rg.Use()
	{
		rg.GET("", cluster.List)
		rg.POST("/create", cluster.Create)
		rg.POST("/:name", cluster.Update)
		rg.DELETE("/:name", cluster.Delete)
	}
}
