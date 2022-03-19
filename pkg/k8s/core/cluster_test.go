package core

import (
	v1 "github.com/gitctl-pro/apps/apis/core/v1"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

func TestCreateCluster(t *testing.T) {
	kubeConfig := "/Users/zsw/.kube/config"
	restConfig, _ := clientcmd.BuildConfigFromFlags("", kubeConfig)
	byteConcfig, _ := ioutil.ReadFile(kubeConfig)
	clusterManager := NewClusterManager(restConfig)
	//err := clusterManager.Delete("dev")
	//if err !=nil {
	//		log.Info(err)
	//	}
	err2 := clusterManager.Create(&v1.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: "dev",
		},
		Spec: v1.ClusterSpec{
			KubeConfig: string(byteConcfig),
		},
	})
	if err2 != nil {
		log.Info(err2)
	}
}
