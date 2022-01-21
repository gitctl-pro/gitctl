package k8s

import "github.com/gin-gonic/gin"

type pod struct{}

func NewPod() PodInterface {
	return &pod{}
}

func (ctl *pod) List(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pod) Containers(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pod) Get(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pod) Events(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pod) PersistentVolumeClaims(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pod) ExecShell(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pod) ExecShellInfo(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *pod) LogDetail(ctx *gin.Context) {
	panic("implement me")
}
