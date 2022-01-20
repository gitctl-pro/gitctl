package routes

import (
	"github.com/gitctl-pro/gitctl/controller/code"
)

func (r *RouteManager) addCodeRoutes(path string) {
	rg := r.gin.Group(path)
	ctl := code.NewController()
	rg.Use()
	{
		rg.GET("/search", ctl.SearchProject)
	}
}
