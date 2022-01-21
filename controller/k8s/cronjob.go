package k8s

import "github.com/gin-gonic/gin"

type cronJob struct{}

func NewCronJob() CronJobInterface {
	return &cronJob{}
}

func (ctl *cronJob) ListCronJob(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *cronJob) GetCronJob(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *cronJob) Events(ctx *gin.Context) {
	panic("implement me")
}
