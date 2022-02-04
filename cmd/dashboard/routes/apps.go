package routes

import (
	"github.com/gitctl-pro/gitctl/controller/apps"
)

func (r *RouteManager) addAppsRoutes(path string) {
	rg := r.gin.Group(path)
	apps := apps.NewController(nil)
	// route: /apps/app
	rg = r.gin.Group(path + "/app")
	rg.Use()
	{
		rg.GET("/:namespace", apps.App.ListApplication)
		rg.POST("/:namespace/:name", apps.App.Create)
		rg.PUT("/:namespace/:name", apps.App.Put)
		rg.DELETE("/namespace/:name", apps.App.Delete)
	}

	// route: /apps/rollout
	rg = r.gin.Group(path + "/rollout")
	rg.Use()
	{
		rg.GET("/:namespace", apps.Rollout.ListRollout)
		rg.POST("/:namespace/:name", apps.Rollout.Create)
		rg.PUT("/:namespace/:name", apps.Rollout.Put)
		rg.DELETE("/namespace/:name", apps.Rollout.Delete)
	}
}
