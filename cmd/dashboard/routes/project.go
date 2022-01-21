package routes

import (
	"github.com/gitctl-pro/gitctl/controller/project"
)

func (r *RouteManager) addProjectRoutes(path string) {
	rg := r.gin.Group(path)
	project := project.NewController()
	rg.Use()
	{
		rg.GET("/issues", project.Issues.ListIssues)
	}
}
