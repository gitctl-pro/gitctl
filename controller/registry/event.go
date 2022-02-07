package registry

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/registry/api"
	"github.com/gitctl-pro/gitctl/pkg/registry/handler"
	"net/http"
)

type event struct {
	Handler handler.Handler
}

func NewEvent() *event {
	return &event{
		Handler: handler.NewHandler(),
	}
}

func (e *event) Record(ctx *gin.Context) {
	recordEvents := &api.RecordEvents{}
	ctx.BindJSON(recordEvents)
	for _, event := range recordEvents.Events {
		switch event.Action {
		case api.ACTION_PUSH:
			if event.Target.Tag == "" {
				break
			}
			go func(event *api.Event) {
				e.Handler.PushAction(event)
			}(event)
		case api.ACTION_PULL:
			go func(event *api.Event) {
				e.Handler.PullAction(event)
			}(event)
		}
	}
	ctx.JSON(http.StatusOK, nil)
}
