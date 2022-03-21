package core

import (
	"context"
	"github.com/gitctl-pro/apps/apis/core/v1"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"github.com/gitctl-pro/gitctl/pkg/k8s/util"
	"github.com/gitctl-pro/gitctl/pkg/k8s/util/controller"
	"github.com/gitctl-pro/gitctl/pkg/logging"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"sync"
	"time"
)

var (
	log = logging.DefaultLogger.WithField("component", "clustermanager")
)

type clusterManager struct {
	clusters map[string]*rest.Config
	resource k8s.Resource
	informer k8s.Informer
	lock     sync.Mutex
}

func NewClusterManager(config *rest.Config) k8s.ClusterManager {
	clusters := make(map[string]*rest.Config, 0)
	resource := k8s.NewResource(config, &schema.GroupVersionKind{
		Kind: "cluster", Group: "core.gitctl.com", Version: "v1",
	})
	informer := k8s.NewInformer(resource, &v1.Cluster{}, 0)
	return &clusterManager{
		clusters: clusters,
		informer: informer,
		resource: resource,
	}
}

func (m *clusterManager) Get(name string) (*rest.Config, error) {
	client, ok := m.clusters[name]
	if !ok {
		return m.Update(name)
	}
	return client, nil
}

func (m *clusterManager) Create(cluster *v1.Cluster) error {
	return m.resource.Create(cluster)
}

func (m *clusterManager) Update(name string) (*rest.Config, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	cluster := &v1.Cluster{}
	err := m.resource.Get(name, cluster)
	if err != nil {
		return nil, err
	}
	config, err := util.LoadConfig(cluster.Spec.KubeConfig)
	m.clusters[name] = config
	return m.clusters[name], nil
}

func (m *clusterManager) Delete(name string) error {
	m.lock.Lock()
	delete(m.clusters, name)
	m.lock.Unlock()
	return m.resource.Delete(name)
}

func (m *clusterManager) Watcher(ctx context.Context) {
	log.Info("clusterManager watcher")
	go m.informer.Run()
	go wait.Until(func() {
		controller.RunWorker(m.informer.Workqueue(), m.resource.Resource(), m.handleWatcher)
	}, time.Second, ctx.Done())
}

func (m *clusterManager) handleWatcher(key string) error {
	_, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	obj := &v1.Cluster{}
	err = m.resource.Get(name, obj)
	if errors.IsNotFound(err) {
		return nil
	}
	m.Update(name)
	return nil
}
