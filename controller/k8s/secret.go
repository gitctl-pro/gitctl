package k8s

import "github.com/gin-gonic/gin"

type secret struct{}

func NewSecret() SecretInterface {
	return &secret{}
}

func (s *secret) ListSecret(ctx *gin.Context) {
	panic("implement me")
}

func (s *secret) GetSecret(ctx *gin.Context) {
	panic("implement me")
}

func (s *secret) CreateSecret(ctx *gin.Context) {
	panic("implement me")
}
