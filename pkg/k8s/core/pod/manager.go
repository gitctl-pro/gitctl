package pod

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
	log = logging.DefaultLogger.WithField("component", "podManager")
)

type podManager struct {
	resource k8s.Resource
	informer k8s.Informer
}

func NewPodManager(config *rest.Config) k8s.PodManager {
	resource := k8s.NewResource(config, &schema.GroupVersionKind{
		Kind: "Pod", Group: "apps", Version: "v1",
	})
	informer := k8s.NewInformer(resource, &v1.Pod{}, 0)
	return &podManager{
		informer: informer,
		resource: resource,
	}
}

func (m *podManager) Watcher(ctx context.Context) {
	log.Info("clusterManager watcher")
	go m.informer.Run(ctx.Done())
	go wait.Until(func() {
		workqueue := m.informer.Workqueue()
		resource := m.resource.Resource()
		controller.RunWorker(workqueue, resource, m.handleWatcher)
	}, time.Second, ctx.Done())
	<-ctx.Done()
}

func (m *podManager) handleWatcher(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	cluster := &v1.Pod{}
	err = m.resource.Namespace(namespace).Get(name, cluster)
	if errors.IsNotFound(err) {
		return nil
	}

	return nil
}
