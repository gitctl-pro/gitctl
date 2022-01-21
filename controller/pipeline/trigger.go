package pipeline

import "github.com/gin-gonic/gin"

type triggerCtl struct{}

func NewTrigger() Trigger {
	return &triggerCtl{}
}

func (t *triggerCtl) GetTrigger(ctx *gin.Context) {
	panic("implement me")
}

func (t *triggerCtl) CreateTrigger(ctx *gin.Context) {
	panic("implement me")
}

func (t *triggerCtl) UpdateTrigger(ctx *gin.Context) {
	panic("implement me")
}

func (t *triggerCtl) ListTrigger(ctx *gin.Context) {
	panic("implement me")
}

func (t *triggerCtl) DeleteTrigger(ctx *gin.Context) {
	panic("implement me")
}
