package code

import "github.com/gin-gonic/gin"

type Registry struct {
	Repo   Repo
	Tag    Tag
	Branch Branch
	Commit Commit
	Tree   Tree
}

func NewController() *Registry {
	return &Registry{
		Repo:   NewRepo(),
		Tag:    NewTag(),
		Branch: NewBranch(),
		Commit: NewCommit(),
		Tree:   NewTree(),
	}
}

type Repo interface {
	GetRepo(ctx *gin.Context)
	UpdateRepo(ctx *gin.Context)
	CreateRepo(ctx *gin.Context)
	ListRepos(ctx *gin.Context)
	DeleteRepo(ctx *gin.Context)
}

type Tag interface {
	GetTag(ctx *gin.Context)
	UpdateTag(ctx *gin.Context)
	CreateTag(ctx *gin.Context)
	ListTags(ctx *gin.Context)
	DeleteTag(ctx *gin.Context)
}

type Branch interface {
	GetBranch(ctx *gin.Context)
	UpdateBranch(ctx *gin.Context)
	CreateBranch(ctx *gin.Context)
	ListBranches(ctx *gin.Context)
	DeleteBranch(ctx *gin.Context)
}

type Commit interface {
	GetCommit(ctx *gin.Context)
	ListCommits(ctx *gin.Context)
}

type Tree interface {
	GetFile(ctx *gin.Context)
	ListPath(ctx *gin.Context)
}
