package pipeline

import (
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

type taskRun struct {
	config *rest.Config
	gvk    *schema.GroupVersionKind
}

func NewTaskRun(config *rest.Config) *taskRun {
	gvk := &schema.GroupVersionKind{
		Group:   "tekton.dev",
		Kind:    "TaskRun",
		Version: "v1beta1",
	}
	return &taskRun{config: config, gvk: gvk}
}

func (t *taskRun) ListTaskRun(ctx *gin.Context) {
	panic("implement me")
}

func (t *taskRun) Get(ctx *gin.Context) {
	panic("implement me")
}

func (t *taskRun) Create(ctx *gin.Context) {
	panic("implement me")
}

func (t *taskRun) Put(ctx *gin.Context) {
	panic("implement me")
}

func (t *taskRun) Delete(ctx *gin.Context) {
	panic("implement me")
}

func (t *taskRun) Log(ctx *gin.Context) {
	panic("implement me")
}
