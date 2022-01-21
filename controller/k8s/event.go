package k8s

import "github.com/gin-gonic/gin"

type event struct{}

func NewEvent() EventInterface {
	return &event{}
}

func (ctl *event) ListEvents(ctx *gin.Context) {
	panic("implement me")
}
