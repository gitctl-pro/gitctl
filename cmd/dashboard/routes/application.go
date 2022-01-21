package routes

import (
	"github.com/gitctl-pro/gitctl/controller/application"
)

func (r *RouteManager) addApplicationRoutes(path string) {
	rg := r.gin.Group(path)
	application := application.NewController()
	// route: /api/v1/cluster
	rg = r.gin.Group(path + "/application")
	rg.Use()
	{
		rg.GET("", application.List)
		rg.POST("/create", application.Create)
		rg.POST("/:namespace/:name", application.Update)
		rg.DELETE("/namespace/:name", application.Delete)
	}
}
