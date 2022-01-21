package k8s

import "github.com/gin-gonic/gin"

type job struct{}

func NewJob() JobInterface {
	return &job{}
}

func (ctl *job) ListJob(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *job) GetJob(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *job) Events(ctx *gin.Context) {
	panic("implement me")
}
