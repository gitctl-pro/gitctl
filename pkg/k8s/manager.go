package k8s

import (
	"context"
	"github.com/gitctl-pro/apps/apis/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type ClusterManager interface {
	Get(name string) (*rest.Config, error)
	Create(cluster *v1.Cluster) error
	Delete(name string) error
}

type Metadata struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type Resource interface {
	Namespace(namespace string) *resource
	Get(name string, object runtime.Object) error
	Put(name string, object runtime.Object) error
	UpdateStatus(name string, object runtime.Object, opts metav1.UpdateOptions) error
	Delete(name string) error
	Create(object runtime.Object) error
	List(object runtime.Object, opts metav1.ListOptions) error
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) error
	MergePatch(name string, patchObject *MergePatchObject) (err error)
	PatchPath(name string, patchObject []PatchPathValue) (err error)
}

type ScaleResource interface {
	Namespace(namespace string) ScaleResource
	ScaleRelicas(name string, replicas int) (err error)
}

type MetaResource interface {
	Namespace(namespace string) MetaResource
	Replace(name string, metadata *Metadata) (err error)
	AddLabel(name string, label, value string) (err error)
	RemoveLabel(name string, label string) (err error)
	AddAnnotation(name string, ann, value string) (err error)
	RemoveAnnotation(name string, ann string) (err error)
}
