package k8s

import "github.com/gin-gonic/gin"

type deploymentCtl struct{}

func NewDeployment() Deployment {
	return &deploymentCtl{}
}

func (ctl *deploymentCtl) ListDeployment(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deploymentCtl) GetEvents(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deploymentCtl) GetDeployment(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deploymentCtl) GetReplicaSets(ctx *gin.Context) {
	panic("implement me")
}
