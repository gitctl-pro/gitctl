package project

import "github.com/gin-gonic/gin"

type issueCtl struct{}

func NewIssue() Issue {
	return &issueCtl{}
}

func (ctl *issueCtl) GetIssue(ctx *gin.Context) {

}

func (ctl *issueCtl) ListIssues(ctx *gin.Context) {

}

func (ctl *issueCtl) UpdateIssue(ctx *gin.Context) {

}

func (ctl *issueCtl) CreateIssue(ctx *gin.Context) {

}
