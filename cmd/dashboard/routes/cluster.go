package routes

import (
	"github.com/gitctl-pro/gitctl/controller/cluster"
)

func (r *RouteManager) addClusterRoutes(path string) {
	rg := r.gin.Group(path)
	cluster := cluster.NewController(r.kubeConfig)
	// route: /cluster
	rg = r.gin.Group(path)
	rg.Use()
	{
		rg.GET("", cluster.List)
		rg.POST("/:name", cluster.Create)
		rg.PUT("/:name", cluster.Put)
		rg.DELETE("/:name", cluster.Delete)
	}
}
