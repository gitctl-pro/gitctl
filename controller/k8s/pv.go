package k8s

import "github.com/gin-gonic/gin"

type pv struct{}

func NewPV() PVInterface {
	return &pv{}
}

func (ctl *pv) DeletePV(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pv) ListPV(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pv) GetPV(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pv) CreatePV(ctx *gin.Context) {
	panic("implement me")
}
