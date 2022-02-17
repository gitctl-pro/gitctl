package istio

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
)

type istio struct {
	VirtualService virtualServiceInterface
}

func NewController(clusterManager k8s.ClusterManager) *istio {
	return &istio{
		VirtualService: NewVirtualService(clusterManager),
	}
}

type virtualServiceInterface interface {
	ListVirtualService(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}
