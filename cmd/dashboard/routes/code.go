package routes

import (
	"github.com/gitctl-pro/gitctl/controller/code"
)

func (r *RouteManager) addCodeRoutes(path string) {
	rg := r.gin.Group(path)
	code := code.NewController()
	rg.Use()
	{
		rg.GET("/repos", code.Repo.ListRepos)
		rg.GET("/:namespace/:name", code.Repo.GetRepo)
		rg.GET("/:namespace/:name/branches", code.Branch.ListBranches)
		rg.GET("/:namespace/:name/tags", code.Tag.ListTags)
		//rg.GET("/:namespace/:name/mergRequest", )
		rg.GET("/:namespace/:name/:branch/commits", code.Commit.ListCommits)
		rg.GET("/:namespace/:name/commit/:commit", code.Commit.GetCommit)
		rg.GET("/:namespace/:name/tree/:branch/:path", code.Tree.GetFile)
	}
}
