package registry

import "github.com/gin-gonic/gin"

type repoCtl struct{}

func NewRepo() Repo {
	return &repoCtl{}
}

func (ctl *repoCtl) GetRepo(ctx *gin.Context) {

}

func (ctl *repoCtl) ListRepos(ctx *gin.Context) {

}

func (ctl *repoCtl) DeleteRepo(ctx *gin.Context) {

}
