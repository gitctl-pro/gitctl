package k8s

import "github.com/gin-gonic/gin"

type namespace struct{}

func NewNamespace() NamespaceInterface {
	return &namespace{}
}

func (ctl *namespace) List(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *namespace) Update(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *namespace) Delete(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *namespace) Create(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *namespace) Events(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *namespace) Quota(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *namespace) LimitRange(ctx *gin.Context) {
	panic("implement me")
}
