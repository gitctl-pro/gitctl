package routes

import (
	"github.com/gitctl-pro/gitctl/controller/code"
)

func (r *RouteManager) addCodeRoutes(path string) {
	rg := r.gin.Group(path)
	ctl := code.NewController()
	rg.Use()
	{
		rg.GET("/repos", ctl.ListRepo)
		rg.GET("/:namespace/:name", ctl.GetRepo)
		rg.GET("/:namespace/:name/branches", ctl.ListBranch)
		rg.GET("/:namespace/:name/tags", ctl.ListTag)
		rg.GET("/:namespace/:name/mergRequest", ctl.ListTag)
		rg.GET("/:namespace/:name/:branch/commits", ctl.ListCommit)
		rg.GET("/:namespace/:name/commit/:commit", ctl.GetCommit)
		rg.GET("/:namespace/:name/tree/:branch/:path", ctl.GetFile)
	}
}
