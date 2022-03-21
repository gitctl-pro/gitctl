package k8s

import (
	"context"
	"github.com/gobuffalo/flect"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"strings"
	"time"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)

type resource struct {
	schema.GroupVersionKind
	client         rest.Interface
	clusterManager ClusterManager
	namespace      string
	resource       string
}

func NewResource(config *rest.Config, gvk *schema.GroupVersionKind) *resource {
	client, resourceName, _ := KindForResource(config, gvk)
	return &resource{
		client:   client,
		resource: resourceName,
	}
}

func KindForResource(config *rest.Config, gvk *schema.GroupVersionKind) (rest.Interface, string, error) {
	config.GroupVersion = &schema.GroupVersion{
		Group:   gvk.Group,
		Version: gvk.Version,
	}
	if len(gvk.Group) == 0 {
		config.APIPath = "/api"
	} else {
		config.APIPath = "/apis"
	}
	config.NegotiatedSerializer = Codecs.WithoutConversion()
	client, err := rest.RESTClientFor(config)
	plural := flect.Pluralize(strings.ToLower(gvk.Kind))
	return client, plural, err
}

func (resource *resource) Namespace(namespace string) Resource {
	resource.namespace = namespace
	return resource
}

func (resource *resource) Resource() string {
	return resource.resource
}

func (resource *resource) Client() rest.Interface {
	return resource.client
}

func (resource *resource) Delete(name string) error {
	defaultPropagationPolicy := metav1.DeletePropagationForeground
	defaultDeleteOptions := &metav1.DeleteOptions{
		PropagationPolicy: &defaultPropagationPolicy,
	}
	req := resource.client.Get().
		Namespace(name).
		Resource(resource.resource).
		Name(name).
		Body(defaultDeleteOptions)

	if len(resource.namespace) > 0 {
		req = req.Namespace(resource.namespace)
	}

	return req.Do(context.TODO()).Error()
}

func (resource *resource) Put(name string, object runtime.Object) (err error) {
	req := resource.client.Put().
		Resource(resource.resource).
		Name(name).
		SetHeader("Content-Type", "application/json").
		Body(object)

	if len(resource.namespace) > 0 {
		req = req.Namespace(resource.namespace)
	}
	err = req.Do(context.TODO()).Error()
	return
}

func (resource *resource) Get(name string, object runtime.Object) (err error) {
	req := resource.client.Get().
		Resource(resource.resource).
		Name(name).
		SetHeader("Accept", "application/json")

	if len(resource.namespace) > 0 {
		req = req.Namespace(resource.namespace)
	}
	err = req.Do(context.TODO()).Into(object)
	return err
}

func (resource *resource) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (err error) {
	req := resource.client.Patch(pt).
		Resource(resource.resource).
		Name(name).
		SubResource(subresources...).
		SetHeader("Accept", "application/json").
		Body(data)

	if len(resource.namespace) > 0 {
		req = req.Namespace(resource.namespace)
	}

	err = req.Do(context.TODO()).Error()
	return err
}

func (resource *resource) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	req := resource.client.Get().
		Resource(resource.resource).
		Timeout(timeout).
		SetHeader("Accept", "application/json")

	if len(resource.namespace) > 0 {
		req = req.Namespace(resource.namespace)
	}
	return req.Watch(ctx)

}

func (resource *resource) Create(object runtime.Object) (err error) {
	req := resource.client.Post().
		Resource(resource.resource).
		SetHeader("Accept", "application/json").
		Body(object)

	if len(resource.namespace) > 0 {
		req = req.Namespace(resource.namespace)
	}
	err = req.Do(context.TODO()).Into(object)
	return err
}

func (resource *resource) List(object runtime.Object, opts metav1.ListOptions) (err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}

	req := resource.client.Get().
		Resource(resource.resource).
		SetHeader("Accept", "application/json").
		Timeout(timeout)

	if len(resource.namespace) > 0 {
		req = req.Namespace(resource.namespace)
	}
	err = req.Do(context.TODO()).Into(object)
	return err
}

func (resource *resource) UpdateStatus(name string, object runtime.Object, opts metav1.UpdateOptions) (err error) {
	req := resource.client.Put().
		Resource("clusters").
		Name(name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec)

	if len(resource.namespace) > 0 {
		req = req.Namespace(resource.namespace)
	}
	err = req.Do(context.TODO()).Into(object)
	return
}
