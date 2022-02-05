package k8s

import (
	gitctl_corev1 "github.com/gitctl-pro/apps/apis/core/v1"
	"github.com/stretchr/testify/assert"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

var (
	kubeConfig = "/Users/zsw/.kube/config"
	cfg, _     = clientcmd.BuildConfigFromFlags("", kubeConfig)
)

func TestResourceCluster(t *testing.T) {
	clusterManager := NewClusterManager(cfg)
	config, _ := clusterManager.Get("dev")
	resource := NewResource(config, &schema.GroupVersionKind{
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
	assert.Equal(t, cluster.Name, "dev")

	resource = NewResource(config, &schema.GroupVersionKind{
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
	assert.Equal(t, deployment.Name, "demo")
}

func TestResourcePod(t *testing.T) {
	resource := NewResource(cfg, &schema.GroupVersionKind{
		Kind:    "pod",
		Version: "v1",
	})
	list := &corev1.PodList{}
	err := resource.Namespace("default").List(list, metav1.ListOptions{})
	if err != nil {
		log.Error(err)
	}
	for k, v := range list.Items {
		log.Info(k, v.Name)
	}
}

func TestResourceNode(t *testing.T) {
	nodeResource := NewResource(cfg, &schema.GroupVersionKind{
		Kind:    "node",
		Group:   "",
		Version: "v1",
	})
	nodeList := &corev1.NodeList{}
	nodeResource.List(nodeList, metav1.ListOptions{})
	for k, v := range nodeList.Items {
		log.Info(k, v.Name)
	}
	assert.Greater(t, len(nodeList.Items), 0)
}

func TestResourceNamespace(t *testing.T) {
	namespaceResource := NewResource(cfg, &schema.GroupVersionKind{
		Kind:    "namespace",
		Group:   "",
		Version: "v1",
	})
	namespaceList := &corev1.NamespaceList{}
	err := namespaceResource.List(namespaceList, metav1.ListOptions{})
	if err != nil {
		log.Error(err)
	}
	for k, v := range namespaceList.Items {
		log.Info(k, v.Name)
	}
	assert.Greater(t, len(namespaceList.Items), 0)
}

func TestResourceCRD(t *testing.T) {
	extResource := NewResource(cfg, &schema.GroupVersionKind{
		Kind:    "CustomResourceDefinition",
		Group:   "apiextensions.k8s.io",
		Version: "v1",
	})
	crdList := &extv1.CustomResourceDefinitionList{}
	extResource.List(crdList, metav1.ListOptions{})
	for k, v := range crdList.Items {
		log.Info(k, v.Name)
	}
	assert.Greater(t, len(crdList.Items), 0)
}

func TestResourceIngress(t *testing.T) {
	resource := NewResource(cfg, &schema.GroupVersionKind{
		Group:   "networking.k8s.io",
		Kind:    "ingress",
		Version: "v1",
	})
	list := &networkingv1.IngressList{}
	err := resource.Namespace("default").List(list, metav1.ListOptions{})
	if err != nil {
		log.Error(err)
	}
	for k, v := range list.Items {
		log.Info(k, v.Name)
	}
}
