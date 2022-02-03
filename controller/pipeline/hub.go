package pipeline

import "github.com/gin-gonic/gin"

type hub struct{}

func NewHub() Hub {
	return &hub{}
}

func (ctl *hub) ListHubs(ctx *gin.Context) {

}
