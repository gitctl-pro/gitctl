package pipeline

import "github.com/gin-gonic/gin"

type history struct{}

func NewHistory() *history {
	return &history{}
}

func (ctl *history) ListHistories(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *history) Get(ctx *gin.Context) {

}

func (ctl *history) Delete(ctx *gin.Context) {
	panic("implement me")
}
