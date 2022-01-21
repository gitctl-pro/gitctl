package k8s

import "github.com/gin-gonic/gin"

type namespace struct{}

func NewNamespace() NamespaceInterface {
	return &namespace{}
}

func (ctl *namespace) UpdateNamespace(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *namespace) DeleteNamespace(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *namespace) CreateNamespace(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *namespace) ListNamespace(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *namespace) GetNamespace(ctx *gin.Context) {
	panic("implement me")
}
