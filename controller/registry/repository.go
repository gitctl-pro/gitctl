package registry

import "github.com/gin-gonic/gin"

type repo struct{}

func NewRepo() *repo {
	return &repo{}
}

func (ctl *repo) UpdateRepo(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *repo) Tags(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *repo) GetRepo(ctx *gin.Context) {

}

func (ctl *repo) ListRepos(ctx *gin.Context) {

}

func (ctl *repo) DeleteRepo(ctx *gin.Context) {

}
