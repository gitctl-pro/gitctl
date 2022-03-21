package deployment

import (
	"context"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"github.com/gitctl-pro/gitctl/pkg/k8s/util/controller"
	"github.com/gitctl-pro/gitctl/pkg/logging"
	"k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"time"
)

var (
	log = logging.DefaultLogger.WithField("component", "deploymentManager")
)

type deploymentManager struct {
	config    *rest.Config
	namespace string
	resource  k8s.Resource
	informer  k8s.Informer
}

func NewDeploymentManager(config *rest.Config) k8s.DeploymentManager {
	resource := k8s.NewResource(config, &schema.GroupVersionKind{
		Kind: "Deployment", Group: "apps", Version: "v1",
	})
	informer := k8s.NewInformer(resource, &v1.Deployment{}, 0)
	return &deploymentManager{
		config:   config,
		informer: informer,
		resource: resource,
	}
}

func (m *deploymentManager) Namespace(namespace string) k8s.DeploymentManager {
	m.namespace = namespace
	return m
}

func (m *deploymentManager) Scale(name string, replicas int) error {
	return k8s.NewScale(m.resource).Namespace(m.namespace).ScaleRelicas(name, replicas)
}

func (m *deploymentManager) Watcher(ctx context.Context) {
	log.Info("deployment watcher")
	go m.informer.Run(ctx.Done())
	go wait.Until(func() {
		workqueue := m.informer.Workqueue()
		resource := m.resource.Resource()
		controller.RunWorker(workqueue, resource, m.handleWatcher)
	}, time.Second, ctx.Done())
	<-ctx.Done()
}

func (m *deploymentManager) handleWatcher(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	obj, exits, err := m.informer.Store().GetByKey(key)
	if errors.IsNotFound(err) || exits {
		log.Info("")
		return nil
	}
	log.Info(namespace, name, obj)
	return nil
}
