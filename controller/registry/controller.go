package registry

import "github.com/gin-gonic/gin"

type RegistryController struct {
	Repo  RepoInterface
	Event EventInterface
}

func NewController() *RegistryController {
	return &RegistryController{
		Repo:  NewRepo(),
		Event: NewEvent(),
	}
}

type RepoInterface interface {
	GetRepo(ctx *gin.Context)
	UpdateRepo(ctx *gin.Context)
	ListRepos(ctx *gin.Context)
	DeleteRepo(ctx *gin.Context)
	Tags(ctx *gin.Context)
}

type TagInterface interface {
	Count(ctx *gin.Context)
	DeleteTag(ctx *gin.Context)
}

type EventInterface interface {
	Record(ctx *gin.Context)
}
