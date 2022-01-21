package pipeline

import "github.com/gin-gonic/gin"

type historyCtl struct{}

func NewHistory() History {
	return &historyCtl{}
}

func (ctl *historyCtl) GetHistory(ctx *gin.Context) {

}

func (ctl *historyCtl) ListHistories(ctx *gin.Context) {

}

func (ctl *historyCtl) Delete(ctx *gin.Context) {

}

func (ctl *historyCtl) DeleteHistory(ctx *gin.Context) {
	panic("implement me")
}
