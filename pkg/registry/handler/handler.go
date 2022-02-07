package handler

import "github.com/gitctl-pro/gitctl/pkg/registry/api"

type Handler interface {
	PushAction(event *api.Event)
	PullAction(event *api.Event)
}

type handler struct{}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) PushAction(event *api.Event) {

}

func (h *handler) PullAction(event *api.Event) {

}
