package k8s

import "github.com/gin-gonic/gin"

type job struct{}

func NewJob() JobInterface {
	return &job{}
}

func (ctl *job) Update(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *job) Delete(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *job) ListJob(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *job) Get(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *job) Events(ctx *gin.Context) {
	panic("implement me")
}
