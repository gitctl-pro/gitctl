package pipeline

import (
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

type trigger struct {
	config *rest.Config
	gvk    *schema.GroupVersionKind
}

func NewTrigger(config *rest.Config) *trigger {
	gvk := &schema.GroupVersionKind{
		Group:   "tekton.dev",
		Kind:    "Trigger",
		Version: "v1beta1",
	}
	return &trigger{config: config, gvk: gvk}
}

func (t *trigger) Get(ctx *gin.Context) {
	panic("implement me")
}

func (t *trigger) Create(ctx *gin.Context) {
	panic("implement me")
}

func (t *trigger) Put(ctx *gin.Context) {
	panic("implement me")
}

func (t *trigger) ListTrigger(ctx *gin.Context) {
	panic("implement me")
}

func (t *trigger) Delete(ctx *gin.Context) {
	panic("implement me")
}

func (t *trigger) Log(ctx *gin.Context) {
	panic("implement me")
}
