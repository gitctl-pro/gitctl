package core

import (
	"context"
	gitctl_corev1 "github.com/gitctl-pro/apps/apis/core/v1"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"github.com/gitctl-pro/gitctl/pkg/k8s/informer"
	"github.com/gitctl-pro/gitctl/pkg/k8s/util"
	controllerutil "github.com/gitctl-pro/gitctl/pkg/k8s/util/controller"
	"github.com/gitctl-pro/gitctl/pkg/logging"
	"k8s.io/apimachinery/pkg/api/errors"
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
	informer *informer.ClusterWatcher
	lock     sync.Mutex
}

func NewClusterManager(config *rest.Config) k8s.ClusterManager {
	clusters := make(map[string]*rest.Config, 0)
	informer := informer.NewClusterWatcher(config)
	return &clusterManager{
		clusters: clusters,
		informer: informer,
		resource: informer.Resource,
	}
}

func (m *clusterManager) Get(name string) (*rest.Config, error) {
	client, ok := m.clusters[name]
	if !ok {
		return m.Update(name)
	}
	return client, nil
}

func (m *clusterManager) Create(cluster *gitctl_corev1.Cluster) error {
	return m.resource.Create(cluster)
}

func (m *clusterManager) Update(name string) (*rest.Config, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	cluster := &gitctl_corev1.Cluster{}
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
		controllerutil.RunWorker(m.informer.Workqueue, "application", m.handleWatcher)
	}, time.Second, ctx.Done())
}

func (m *clusterManager) handleWatcher(key string) error {
	_, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	cluster := &gitctl_corev1.Cluster{}
	err = m.resource.Get(name, cluster)
	if errors.IsNotFound(err) {
		return nil
	}
	m.Update(name)
	return nil
}
