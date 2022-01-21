package code

import "github.com/gin-gonic/gin"

type branchCtl struct{}

func NewBranch() Branch {
	return &branchCtl{}
}

func (ctl *branchCtl) GetBranch(ctx *gin.Context) {

}

func (ctl *branchCtl) ListBranches(ctx *gin.Context) {

}

func (ctl *branchCtl) DeleteBranch(ctx *gin.Context) {

}

func (ctl *branchCtl) UpdateBranch(ctx *gin.Context) {

}

func (ctl *branchCtl) CreateBranch(ctx *gin.Context) {

}
