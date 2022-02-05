package setup

import (
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func SetupK8s(kubeconfig string) (k8s.ClusterManager, *rest.Config, error) {
	log.Info("SetupK8sClient...")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	clusterManager := k8s.NewClusterManager(config)
	return clusterManager, config, nil
}
