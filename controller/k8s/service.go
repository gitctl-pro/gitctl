package k8s

import "github.com/gin-gonic/gin"

type service struct{}

func NewService() ServiceInterface {
	return &service{}
}

func (ctl *service) ListService(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *service) GetService(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *service) Events(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *service) Pods(ctx *gin.Context) {
	panic("implement me")
}
