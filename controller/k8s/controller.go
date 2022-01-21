package k8s

import "github.com/gin-gonic/gin"

type K8s struct {
	Deployment Deployment
}

func NewController() *K8s {
	return &K8s{
		Deployment: NewDeployment(),
	}
}

type Deployment interface {
	GetEvents(ctx *gin.Context)
	GetDeployment(ctx *gin.Context)
	ListDeployment(ctx *gin.Context)
	GetReplicaSets(ctx *gin.Context)
}

type Pod interface {
	ListContainers(ctx *gin.Context)
}
