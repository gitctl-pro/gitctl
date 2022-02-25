package Meta

import (
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

type MetaResource struct {
	resource k8s.Resource
	client   rest.Interface
}

func NewMetaResource(cfg *rest.Config, gvk *schema.GroupVersionKind) k8s.MetaResource {
	resource := k8s.NewResource(cfg, gvk)
	return &MetaResource{
		client:   resource.Client(),
		resource: resource,
	}
}

func (r *MetaResource) Namespace(namespace string) k8s.MetaResource {
	r.resource.Namespace(namespace)
	return r
}

func (r *MetaResource) AddLabel(name string, label, value string) (err error) {
	patchObject := []k8s.PatchPathValue{{
		Op:    "add",
		Path:  "/metadata/labels/" + label,
		Value: value,
	}}
	return r.resource.PatchPath(name, patchObject)
}

func (r *MetaResource) RemoveLabel(name string, label string) (err error) {
	patchObject := []k8s.PatchPathValue{{
		Op:    "remove",
		Path:  "/metadata/labels/" + label,
		Value: "",
	}}
	return r.resource.PatchPath(name, patchObject)
}

func (r *MetaResource) AddAnnotation(name string, annotation, value string) (err error) {
	patchObject := []k8s.PatchPathValue{{
		Op:    "add",
		Path:  "/metadata/annotations/" + annotation,
		Value: value,
	}}
	return r.resource.PatchPath(name, patchObject)
}

func (r *MetaResource) RemoveAnnotation(name string, annotation string) (err error) {
	patchObject := []k8s.PatchPathValue{{
		Op:    "remove",
		Path:  "/metadata/annotations/" + annotation,
		Value: "",
	}}
	return r.resource.PatchPath(name, patchObject)
}

func (r *MetaResource) Replace(name string, meta *k8s.Metadata) (err error) {
	return r.resource.MergePatch(name, &k8s.MergePatchObject{
		Metadata: meta,
	})
}
