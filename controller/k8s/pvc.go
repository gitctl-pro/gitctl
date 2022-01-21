package k8s

import "github.com/gin-gonic/gin"

type pvc struct{}

func NewPVC() PVCInterface {
	return &pvc{}
}

func (ctl *pvc) ListPVC(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pvc) GetPVC(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pvc) CreatePVC(ctx *gin.Context) {
	panic("implement me")
}
