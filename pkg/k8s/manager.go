package k8s

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type ClusterManager interface {
	Get(name string) (rest.Interface, rest.Config)
}

type VerberManager interface {
	Create(object *runtime.Unknown) (runtime.Object, error)
	Put(name string, object *runtime.Unknown) error
	Delete(name string) error
	List(opts metav1.ListOptions) (runtime.Object, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (runtime.Object, error)
}
