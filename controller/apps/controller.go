package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
)

type apps struct {
	App     ApplicationInterface
	Rollout RolloutInterface
}

func NewController(clusterManager k8s.ClusterManager) *apps {
	return &apps{
		App:     NewApplication(clusterManager),
		Rollout: NewRollout(clusterManager),
	}
}

type ApplicationInterface interface {
	ListApplication(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type RolloutInterface interface {
	ListRollout(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}
