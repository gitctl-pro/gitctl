package k8s

import "github.com/gin-gonic/gin"

type replicaSet struct{}

func NewReplicaSet() ReplicaSetInterface {
	return &replicaSet{}
}

func (ctl *replicaSet) List(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *replicaSet) Get(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *replicaSet) Events(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *replicaSet) Pods(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *replicaSet) Service(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *replicaSet) Delete(ctx *gin.Context) {
	panic("implement me")
}
