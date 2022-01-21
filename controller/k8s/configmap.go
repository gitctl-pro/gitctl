package k8s

import "github.com/gin-gonic/gin"

type configMap struct{}

func NewConfigMap() ConfigMapInterface {
	return &configMap{}
}

func (c configMap) ListConfigMap(ctx *gin.Context) {
	panic("implement me")
}

func (c configMap) Get(ctx *gin.Context) {
	panic("implement me")
}

func (c configMap) Create(ctx *gin.Context) {
	panic("implement me")
}

func (c configMap) Update(ctx *gin.Context) {
	panic("implement me")
}

func (c configMap) Delete(ctx *gin.Context) {
	panic("implement me")
}
