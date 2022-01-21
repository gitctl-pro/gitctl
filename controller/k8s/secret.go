package k8s

import "github.com/gin-gonic/gin"

type secret struct{}

func NewSecret() SecretInterface {
	return &secret{}
}

func (s *secret) Delete(ctx *gin.Context) {
	panic("implement me")
}

func (s *secret) ListSecret(ctx *gin.Context) {
	panic("implement me")
}

func (s *secret) Get(ctx *gin.Context) {
	panic("implement me")
}

func (s *secret) Create(ctx *gin.Context) {
	panic("implement me")
}
