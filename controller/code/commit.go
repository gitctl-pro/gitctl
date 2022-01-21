package code

import "github.com/gin-gonic/gin"

type commitCtl struct{}

func NewCommit() Commit {
	return &commitCtl{}
}

func (ctl *commitCtl) GetCommit(ctx *gin.Context) {

}

func (ctl *commitCtl) ListCommits(ctx *gin.Context) {

}
