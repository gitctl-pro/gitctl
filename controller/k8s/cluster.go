package k8s

import "github.com/gin-gonic/gin"

type cluster struct{}

func NewCluster() ClusterInterface {
	return &cluster{}
}

func (ctl *cluster) ListCluster(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *cluster) GetCluster(ctx *gin.Context) {
	panic("implement me")
}
