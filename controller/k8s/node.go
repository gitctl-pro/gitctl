package k8s

import "github.com/gin-gonic/gin"

type node struct{}

func NewNode() NodeInterface {
	return &node{}
}

func (ctl *node) ListNode(ctx *gin.Context) {
	panic("implement me")
}

func (ctl *node) GetNode(ctx *gin.Context) {
	panic("implement me")
}
