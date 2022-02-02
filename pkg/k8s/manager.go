package k8s

import (
	"context"
	"github.com/gitctl-pro/apps/apis/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
)

type ClusterManager interface {
	Get(name string) (*kubeClient, error)
	Create(cluster *v1.Cluster) error
	Delete(name string) error
}

type Resource interface {
	Get(name string, object runtime.Object) error
	Put(name string, object runtime.Object) error
	Delete(name string) error
	Create(object runtime.Object) error
	List(object runtime.Object, opts metav1.ListOptions) error
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) error
}
