package routes

import (
	"github.com/gitctl-pro/gitctl/controller/k8s"
)

func (r *RouteManager) addK8sRoutes(path string) {
	rg := r.gin.Group(path)
	k8s := k8s.NewController()
	rg.Use()
	{
		rg.GET("/deployment/:namespace", k8s.Deployment.ListDeployment)
		rg.GET("/deployment/:namespace/:name", k8s.Deployment.GetDeployment)
	}
}
