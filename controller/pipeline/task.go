package pipeline

import "github.com/gin-gonic/gin"

type taskCtl struct{}

func NewTask() Task {
	return &taskCtl{}
}

func (t *taskCtl) GetTask(ctx *gin.Context) {
	panic("implement me")
}

func (t *taskCtl) CreateTask(ctx *gin.Context) {
	panic("implement me")
}

func (t *taskCtl) UpdateTask(ctx *gin.Context) {
	panic("implement me")
}

func (t *taskCtl) ListTask(ctx *gin.Context) {
	panic("implement me")
}

func (t *taskCtl) DeleteTask(ctx *gin.Context) {
	panic("implement me")
}

func (t *taskCtl) Log(ctx *gin.Context) {
	panic("implement me")
}
