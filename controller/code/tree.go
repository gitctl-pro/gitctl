package code

import "github.com/gin-gonic/gin"

type treeCtl struct{}

func NewTree() Tree {
	return &treeCtl{}
}

func (ctl *treeCtl) GetFile(ctx *gin.Context) {

}

func (ctl *treeCtl) ListPath(ctx *gin.Context) {

}
