package k8s

import "github.com/gin-gonic/gin"

type configmap struct{}

func NewConfigmap() ConfigmapInterface {
	return &configmap{}
}

func (ctl *configmap) List(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *configmap) Get(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *configmap) Create(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *configmap) Update(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *configmap) Events(ctx *gin.Context) {
	panic("implement me")
}
