package pipeline

import "github.com/gin-gonic/gin"

type hubCtl struct{}

func NewHub() Hub {
	return &hubCtl{}
}

func (ctl *hubCtl) ListHubs(ctx *gin.Context) {

}
