package cluster

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

func NewController(cfg *rest.Config) *cluster {
	resource := k8s.NewResource(cfg, &schema.GroupVersionKind{
		Kind:    "cluster",
		Group:   "core.gitctl.com",
		Version: "v1",
	})
	return &cluster{config: cfg, resource: resource}
}

type ClusterInterface interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Put(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}
