package registry

import "github.com/gin-gonic/gin"

type RegistryController struct {
	Repo Repo
}

func NewController() *RegistryController {
	return &RegistryController{
		Repo: NewRepo(),
	}
}

type Repo interface {
	GetRepo(ctx *gin.Context)
	UpdateRepo(ctx *gin.Context)
	ListRepos(ctx *gin.Context)
	DeleteRepo(ctx *gin.Context)
	Tags(ctx *gin.Context)
}

type Tag interface {
	Count(ctx *gin.Context)
	DeleteTag(ctx *gin.Context)
}
