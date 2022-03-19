package apps

import (
	"context"
	gitctl_corev1 "github.com/gitctl-pro/apps/apis/core/v1"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"github.com/gitctl-pro/gitctl/pkg/k8s/informer"
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
	log = logging.DefaultLogger.WithField("component", "deploymentManager")
)

type deploymentManager struct {
	Resource k8s.Resource
	informer *informer.DeploymentWatcher
	lock     sync.Mutex
}

func NewDeploymentManager(config *rest.Config) k8s.DeploymentManager {
	informer := informer.NewDeploymentWatcher(config)
	return &deploymentManager{
		informer: informer,
		Resource: informer.Resource,
	}
}

func (m *deploymentManager) Watcher(ctx context.Context) {
	log.Info("clusterManager watcher")
	go m.informer.Run()
	go wait.Until(func() {
		controllerutil.RunWorker(m.informer.Workqueue, "application", m.handleWatcher)
	}, time.Second, ctx.Done())
}

func (m *deploymentManager) handleWatcher(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	cluster := &gitctl_corev1.Cluster{}
	err = m.Resource.Namespace(namespace).Get(name, cluster)
	if errors.IsNotFound(err) {
		return nil
	}

	return nil
}
