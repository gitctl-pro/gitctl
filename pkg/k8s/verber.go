package k8s

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	"time"
)

type resourceVerber struct {
	schema.GroupVersionKind
	client         rest.Interface
	clusterManager ClusterManager
	namespace      string
	resource       string
}

func NewResourceVerber(client rest.Interface, kind schema.GroupVersionKind) *resourceVerber {
	client, resource, _ := KindResource(kind)
	return &resourceVerber{
		client:   client,
		resource: resource,
	}
}

func KindResource(kind schema.GroupVersionKind) (rest.Interface, string, error) {

	return nil, "", nil
}

func (verber *resourceVerber) Cluster(cluster string) {
	verber.client, _ = verber.clusterManager.Get(cluster)
}

func (verber *resourceVerber) Namespace(namespace string) *resourceVerber {
	verber.namespace = namespace
	return verber
}

func (verber *resourceVerber) Delete(name string) error {
	defaultPropagationPolicy := metav1.DeletePropagationForeground
	defaultDeleteOptions := &metav1.DeleteOptions{
		PropagationPolicy: &defaultPropagationPolicy,
	}
	req := verber.client.Get().
		Namespace(name).
		Resource(verber.resource).
		Name(name).
		Body(defaultDeleteOptions)

	if len(verber.namespace) > 0 {
		req.Namespace(verber.namespace)
	}

	return req.Do(context.TODO()).Error()
}

func (verber *resourceVerber) Put(name string, object *runtime.Unknown) error {

	req := verber.client.Put().
		Resource(verber.resource).
		Name(name).
		SetHeader("Content-Type", "application/json").
		Body([]byte(object.Raw))

	if len(verber.namespace) > 0 {
		req.Namespace(verber.namespace)
	}

	return req.Do(context.TODO()).Error()
}

func (verber *resourceVerber) Get(name string) (runtime.Object, error) {
	result := &runtime.Unknown{}
	req := verber.client.Get().
		Resource(verber.resource).
		Name(name).
		SetHeader("Accept", "application/json")

	if len(verber.namespace) > 0 {
		req.Namespace(verber.namespace)
	}

	err := req.Do(context.TODO()).Into(result)
	return result, err
}

func (verber *resourceVerber) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (runtime.Object, error) {
	result := &runtime.Unknown{}
	req := verber.client.Patch(pt).
		Resource(verber.resource).
		Name(name).
		SubResource(subresources...).
		SetHeader("Accept", "application/json").
		Body(data)

	if len(verber.namespace) > 0 {
		req.Namespace(verber.namespace)
	}

	err := req.Do(context.TODO()).Into(result)
	return result, err
}

func (verber *resourceVerber) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	req := verber.client.Get().
		Resource(verber.resource).
		Timeout(timeout).
		SetHeader("Accept", "application/json")

	if len(verber.namespace) > 0 {
		req.Namespace(verber.namespace)
	}
	return req.Watch(ctx)

}

func (verber *resourceVerber) Create(object *runtime.Unknown) (result runtime.Object, err error) {
	req := verber.client.Post().
		Resource(verber.resource).
		SetHeader("Accept", "application/json").
		Body([]byte(object.Raw))

	if len(verber.namespace) > 0 {
		req.Namespace(verber.namespace)
	}
	err = req.Do(context.TODO()).Into(result)
	return result, err
}

func (verber *resourceVerber) List(opts metav1.ListOptions) (result runtime.Object, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}

	req := verber.client.Get().
		Resource(verber.resource).
		SetHeader("Accept", "application/json").
		Timeout(timeout)

	if len(verber.namespace) > 0 {
		req.Namespace(verber.namespace)
	}

	err = req.Do(context.TODO()).Into(result)

	return
}
