package pipeline

import (
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

type task struct {
	config *rest.Config
	gvk    *schema.GroupVersionKind
}

func NewTask(config *rest.Config) *task {
	gvk := &schema.GroupVersionKind{
		Group:   "tekton.dev",
		Kind:    "Task",
		Version: "v1beta1",
	}
	return &task{config: config, gvk: gvk}
}

func (t *task) ListTask(ctx *gin.Context) {
	panic("implement me")
}

func (t *task) Get(ctx *gin.Context) {
	panic("implement me")
}

func (t *task) Create(ctx *gin.Context) {
	panic("implement me")
}

func (t *task) Put(ctx *gin.Context) {
	panic("implement me")
}

func (t *task) Delete(ctx *gin.Context) {
	panic("implement me")
}

func (t *task) Log(ctx *gin.Context) {
	panic("implement me")
}
