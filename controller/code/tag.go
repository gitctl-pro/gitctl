package code

import "github.com/gin-gonic/gin"

type tagCtl struct{}

func NewTag() Tag {
	return &tagCtl{}
}

func (ctl *tagCtl) GetTag(ctx *gin.Context) {

}

func (ctl *tagCtl) ListTags(ctx *gin.Context) {

}

func (ctl *tagCtl) DeleteTag(ctx *gin.Context) {

}

func (ctl *tagCtl) UpdateTag(ctx *gin.Context) {

}

func (ctl *tagCtl) CreateTag(ctx *gin.Context) {

}
