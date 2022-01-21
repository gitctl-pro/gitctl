package project

import "github.com/gin-gonic/gin"

type ProjectController struct {
	Issues Issue
}

func NewController() *ProjectController {
	return &ProjectController{
		Issues: NewIssue(),
	}
}

type Issue interface {
	GetIssue(ctx *gin.Context)
	ListIssues(ctx *gin.Context)
	UpdateIssue(ctx *gin.Context)
	CreateIssue(ctx *gin.Context)
}
