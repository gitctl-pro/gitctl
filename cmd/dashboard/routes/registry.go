package routes

import (
	"github.com/gitctl-pro/gitctl/controller/registry"
)

func (r *RouteManager) addRegistryRoutes(path string) {
	rg := r.gin.Group(path)
	registry := registry.NewController()
	rg.Use()
	{
		rg.GET("/repos", registry.Repo.ListRepos)
	}
}
