package k8s

import (
	gitctl_corev1 "github.com/gitctl-pro/apps/apis/core/v1"
	appv1 "k8s.io/api/apps/v1"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

func TestReourceVerber(t *testing.T) {
	kubeConfig := "/Users/zsw/.kube/config"
	restConfig, _ := clientcmd.BuildConfigFromFlags("", kubeConfig)
	clusterManager := NewClusterManager(restConfig)
	client, _ := clusterManager.Get("dev")
	resource := NewResourceVerber(client.config, &schema.GroupVersionKind{
		Kind:    "cluster",
		Group:   "core.gitctl.com",
		Version: "v1",
	})
	cluster := &gitctl_corev1.Cluster{}
	err := resource.Get("dev", cluster)

	log.Info(cluster.Name)
	if err != nil {
		log.Error(err)
	}

	resource = NewResourceVerber(client.config, &schema.GroupVersionKind{
		Kind:    "Deployment",
		Group:   "apps",
		Version: "v1",
	})
	deployment := &appv1.Deployment{}
	err = resource.Namespace("default").Get("demo", deployment)
	log.Info(deployment.Name)
	if err != nil {
		log.Error(err)
	}

	extResource := NewResourceVerber(client.config, &schema.GroupVersionKind{
		Kind:    "CustomResourceDefinition",
		Group:   "apiextensions.k8s.io",
		Version: "v1",
	})
	crdList := &extv1.CustomResourceDefinitionList{}
	extResource.List(crdList, metav1.ListOptions{})
	for k, v := range crdList.Items {
		log.Info(k, v.Name)
	}
}
