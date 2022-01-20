package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/config"
	"k8s.io/client-go/rest"
)

type RouteManager struct {
	kubeConfig     *rest.Config
	configResolver *config.ConfigResolver
	gin            *gin.Engine
}

func NewRouteManager(configResolver *config.ConfigResolver) *RouteManager {
	gin := gin.Default()
	return &RouteManager{
		configResolver: configResolver,
		gin:            gin,
	}
}

func (r *RouteManager) Run() {
	r.addCodeRoutes("/code")
	r.gin.Run(":8081")
}
