package cluster

import (
	"github.com/gin-gonic/gin"
)

func NewController() ClusterInterface {
	return &cluster{}
}

type ClusterInterface interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}
