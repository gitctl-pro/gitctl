package setup

import (
	"github.com/gitctl-pro/gitctl/pkg/config"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"github.com/gitctl-pro/gitctl/pkg/k8s/core"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func SetupK8s(kubeConfigPath string) (k8s.ClusterManager, *rest.Config, error) {
	log.Info("SetupK8sClient...")
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	clusterManager := core.NewClusterManager(config)
	return clusterManager, config, nil
}

func SetupK8sWacther(config *config.ConfigResolver, kubeConfigPath string) (*k8s.K8sWatcher, *rest.Config) {
	log.Info("SetupK8sClient...")
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	return k8s.NewK8sWatcher(config, kubeConfig), kubeConfig
}
