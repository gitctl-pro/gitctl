package k8s

import "github.com/gin-gonic/gin"

type hpa struct{}

func NewHPA() HPAInterface {
	return &hpa{}
}

func (ctl *hpa) GetHPA(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *hpa) ListHPA(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *hpa) CreateHPA(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *hpa) PutHPA(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *hpa) DeleteHPA(ctx *gin.Context) {
	panic("implement me")
}
