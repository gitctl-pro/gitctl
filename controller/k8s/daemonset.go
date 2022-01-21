package k8s

import "github.com/gin-gonic/gin"

type daemonset struct{}

func NewDaemonset() DaemonsetInterface {
	return &daemonset{}
}

func (ctl *daemonset) ListDaemonset(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *daemonset) GetDaemonset(ctx *gin.Context) {
	panic("implement me")
}
