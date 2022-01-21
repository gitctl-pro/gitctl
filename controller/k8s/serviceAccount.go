package k8s

import "github.com/gin-gonic/gin"

type serviceAccount struct{}

func NewServiceAccount() ServiceAccountInterface {
	return &serviceAccount{}
}

func (s *serviceAccount) ListServiceAccount(ctx *gin.Context) {
	panic("implement me")
}

func (s *serviceAccount) Get(ctx *gin.Context) {
	panic("implement me")
}

func (s *serviceAccount) Delete(ctx *gin.Context) {
	panic("implement me")
}
