package k8s

import "github.com/gin-gonic/gin"

type pvc struct{}

func NewPVC() PVCInterface {
	return &pvc{}
}

func (ctl *pvc) Delete(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pvc) ListPVC(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pvc) Get(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pvc) Create(ctx *gin.Context) {
	panic("implement me")
}
