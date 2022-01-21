package k8s

import "github.com/gin-gonic/gin"

type scale struct{}

func NewScale() ScaleInterface {
	return &scale{}
}

func (ctl *scale) ScaleResource(ctx *gin.Context) {
	panic("implement me")
}
