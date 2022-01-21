package k8s

import "github.com/gin-gonic/gin"

type ingress struct{}

func NewIngress() IngressInterface {
	return &ingress{}
}

func (ctl *ingress) Get(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *ingress) Update(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *ingress) Delete(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *ingress) Events(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *ingress) ListEvent(ctx *gin.Context) {
	panic("implement me")
}
