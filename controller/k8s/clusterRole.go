package k8s

import "github.com/gin-gonic/gin"

type clusterRole struct{}

func NewClusterRole() ClusterRoleInterface {
	return &clusterRole{}
}

func (c clusterRole) ListClusterRole(ctx *gin.Context) {
	panic("implement me")
}

func (c clusterRole) Get(ctx *gin.Context) {
	panic("implement me")
}

func (c clusterRole) Delete(ctx *gin.Context) {
	panic("implement me")
}
