package application

import (
	"github.com/gin-gonic/gin"
)

func NewController() ApplicationInterface {
	return &application{}
}

type ApplicationInterface interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}
