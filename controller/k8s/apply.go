package k8s

import "github.com/gin-gonic/gin"

type apply struct{}

func NewApply() ApplyInterface {
	return &apply{}
}

func (ctl *apply) Apply(ctx *gin.Context) {
	panic("implement me")
}
