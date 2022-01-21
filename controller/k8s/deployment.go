package k8s

import "github.com/gin-gonic/gin"

type deployment struct{}

func NewDeployment() DeploymentInterface {
	return &deployment{}
}

func (ctl *deployment) Events(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) Get(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) List(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) ReplicaSets(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) NewReplicaSets(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) OldReplicaSets(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) RolloutRestart(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) RolloutPause(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) RolloutResume(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *deployment) RolloutRollback(ctx *gin.Context) {
	panic("implement me")
}
