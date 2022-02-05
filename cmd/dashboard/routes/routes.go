package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/config"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/client-go/rest"
)

type RouteManager struct {
	kubeConfig     *rest.Config
	clusterManager k8s.ClusterManager
	configResolver *config.ConfigResolver
	gin            *gin.Engine
}

func NewRouteManager(kubeConfig *rest.Config, clusterManager k8s.ClusterManager, configResolver *config.ConfigResolver) *RouteManager {
	gin := gin.Default()
	return &RouteManager{
		kubeConfig:     kubeConfig,
		clusterManager: clusterManager,
		configResolver: configResolver,
		gin:            gin,
	}
}

func (r *RouteManager) Run() {
	r.addCodeRoutes("/code")
	r.addK8sRoutes("/k8s/:cluster")
	r.addAppsRoutes("/apps")
	r.addClusterRoutes("/cluster")
	r.addPipelineRoutes("/pipeline")
	r.addRegistryRoutes("/registry")
	r.addProjectRoutes("/project")
	r.gin.Run(":8081")
}
