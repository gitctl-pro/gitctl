package event

import (
	"context"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"github.com/gitctl-pro/gitctl/pkg/k8s/util/controller"
	"github.com/gitctl-pro/gitctl/pkg/logging"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"time"
)

var (
	log = logging.DefaultLogger.WithField("component", "eventManager")
)

type eventManager struct {
	Resource k8s.Resource
	informer k8s.Informer
}

func NewEventManager(config *rest.Config) k8s.EventManager {
	resource := k8s.NewResource(config, &schema.GroupVersionKind{
		Kind: "Event", Group: "apps", Version: "v1",
	})
	informer := k8s.NewInformer(resource, &v1.Event{}, 0)
	return &eventManager{
		informer: informer,
		Resource: resource,
	}
}

func (m *eventManager) Watcher(ctx context.Context) {
	log.Info("clusterManager watcher")
	go m.informer.Run(ctx.Done())
	go wait.Until(func() {
		workqueue := m.informer.Workqueue()
		resource := m.Resource.Resource()
		controller.RunWorker(workqueue, resource, m.handleWatcher)
	}, time.Second, ctx.Done())
	<-ctx.Done()
}

func (m *eventManager) handleWatcher(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	obj := &v1.Event{}

	err = m.Resource.Namespace(namespace).Get(name, obj)
	if errors.IsNotFound(err) {
		return nil
	}

	return nil
}
