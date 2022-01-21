package k8s

import "github.com/gin-gonic/gin"

type ingress struct{}

func NewIngress() IngressInterface {
	return &ingress{}
}

func (ctl *ingress) ListIngress(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *ingress) GetIngress(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *ingress) Events(ctx *gin.Context) {
	panic("implement me")
}
