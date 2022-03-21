package watcher

import (
	"context"
	"github.com/gitctl-pro/gitctl/pkg/config"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"github.com/gitctl-pro/gitctl/pkg/k8s/apps/deployment"
	"github.com/gitctl-pro/gitctl/pkg/k8s/core"
	"github.com/gitctl-pro/gitctl/pkg/logging"
	"k8s.io/client-go/rest"
	"os"
	"os/signal"
	"syscall"
)

var (
	log = logging.DefaultLogger.WithField("component", "watcher")
)

type K8sWatcher struct {
	configResolver    *config.ConfigResolver
	kubeConfig        *rest.Config
	clusterManager    k8s.ClusterManager
	deploymentManager k8s.DeploymentManager
	stopChan          chan struct{}
}

func NewK8sWatcher(configResolver *config.ConfigResolver, kubeConfig *rest.Config) *K8sWatcher {
	return &K8sWatcher{
		configResolver:    configResolver,
		kubeConfig:        kubeConfig,
		stopChan:          make(chan struct{}),
		deploymentManager: deployment.NewDeploymentManager(kubeConfig),
		clusterManager:    core.NewClusterManager(kubeConfig),
	}
}

func (w *K8sWatcher) EnableClusterWatcher() k8s.ClusterManager {
	log.Info("enable k8s clusterwatcher")
	w.clusterManager.Watcher(context.Background())
	return w.clusterManager
}

func (w *K8sWatcher) EnableDeploymentWatcher() k8s.ClusterManager {
	log.Info("enable k8s deploywatcher")
	w.deploymentManager.Watcher(context.Background())
	return w.clusterManager
}

func (w *K8sWatcher) EnableEventWatcher() {

}

func (w *K8sWatcher) EnablePodWatcher() {

}

func (w *K8sWatcher) EnableReplicasetWatcher() {

}

func (k *K8sWatcher) Run() {
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, syscall.SIGINT, syscall.SIGTERM)
	<-terminate
	if k.stopChan != nil {
		close(k.stopChan)
	}
}
