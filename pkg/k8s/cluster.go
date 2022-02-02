package k8s

import (
	corev1 "github.com/gitctl-pro/apps/apis/core/v1"
	"github.com/gitctl-pro/gitctl/pkg/k8s/util"
	"github.com/gitctl-pro/gitctl/pkg/logging"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"sync"
)

var (
	log = logging.DefaultLogger.WithField("component", "clustermanager")
)

type kubeClient struct {
	config *rest.Config
}

type clusterManager struct {
	clusters map[string]*kubeClient
	resource *resourceVerber
	lock     sync.Mutex
}

func NewClusterManager(config *rest.Config) ClusterManager {
	clusters := make(map[string]*kubeClient, 0)
	resource := NewResourceVerber(config, &schema.GroupVersionKind{
		Kind:    "cluster",
		Group:   "core.gitctl.com",
		Version: "v1",
	})
	return &clusterManager{
		clusters: clusters,
		resource: resource,
	}
}

func (m *clusterManager) Get(name string) (*kubeClient, error) {
	client, ok := m.clusters[name]
	if !ok {
		m.lock.Lock()
		defer m.lock.Unlock()
		cluster := &corev1.Cluster{}
		err := m.resource.Get(name, cluster)
		if err != nil {
			return nil, err
		}
		config, err := util.LoadConfig(cluster.Spec.KubeConfig)
		m.clusters[name] = &kubeClient{
			config: config,
		}
		return m.clusters[name], nil
	}
	return client, nil
}

func (m *clusterManager) Create(cluster *corev1.Cluster) error {
	return m.resource.Create(cluster)
}

func (m *clusterManager) Delete(name string) error {
	m.lock.Lock()
	delete(m.clusters, name)
	m.lock.Unlock()
	return m.resource.Delete(name)
}
