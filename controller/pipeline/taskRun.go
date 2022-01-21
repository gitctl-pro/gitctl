package pipeline

import "github.com/gin-gonic/gin"

type taskRunCtl struct{}

func NewTaskRun() TaskRun {
	return &taskRunCtl{}
}

func (t *taskRunCtl) GetTaskRun(ctx *gin.Context) {
	panic("implement me")
}

func (t *taskRunCtl) CreateTaskRun(ctx *gin.Context) {
	panic("implement me")
}

func (t *taskRunCtl) ListTaskRun(ctx *gin.Context) {
	panic("implement me")
}

func (t *taskRunCtl) Log(ctx *gin.Context) {
	panic("implement me")
}
