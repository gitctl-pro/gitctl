package k8s

import "github.com/gin-gonic/gin"

type application struct{}

func NewApplication() ApplicationInterface {
	return &application{}
}

func (ctl *application) UpdateApplication(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *application) DeleteApplication(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *application) CreateApplication(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *application) ListApplication(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *application) GetApplication(ctx *gin.Context) {
	panic("implement me")
}
