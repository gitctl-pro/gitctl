package routes

import (
	"github.com/gitctl-pro/gitctl/controller/istio"
)

func (r *RouteManager) addIstioRoutes(path string) {
	rg := r.gin.Group(path)
	istio := istio.NewController(r.clusterManager)
	// route: /istio/virtualService
	rg = r.gin.Group(path + "/virtualService")
	rg.Use()
	{
		rg.GET("/:namespace", istio.VirtualService.ListVirtualService)
		rg.POST("/:namespace/:name", istio.VirtualService.Create)
		rg.PUT("/:namespace/:name", istio.VirtualService.Put)
		rg.DELETE("/namespace/:name", istio.VirtualService.Delete)
	}
}
