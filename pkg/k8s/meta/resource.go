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

type Metadata struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

func NewMetaResource(cfg *rest.Config, gvk *schema.GroupVersionKind) *MetaResource {
	resource := k8s.NewResource(cfg, gvk)
	return &MetaResource{
		client:   resource.Client(),
		resource: resource,
	}
}

func (r *MetaResource) AddLabel(name string, label, value string) (err error) {
	patchObject := []k8s.PatchPathValue{{
		Op:    "add",
		Path:  "/metadata/labels/" + label,
		Value: value,
	}}
	return r.resource.PatchPath(name, patchObject)
}

func (r *MetaResource) RemoveLabel(name string, label, value string) (err error) {
	patchObject := []k8s.PatchPathValue{{
		Op:    "add",
		Path:  "/metadata/labels/" + label,
		Value: value,
	}}
	return r.resource.PatchPath(name, patchObject)
}

func (r *MetaResource) AddAnnotations(name string, annotation, value string) (err error) {
	patchObject := []k8s.PatchPathValue{{
		Op:    "add",
		Path:  "/metadata/annotations/" + annotation,
		Value: value,
	}}
	return r.resource.PatchPath(name, patchObject)
}

func (r *MetaResource) RemoveAnnotations(name string, annotation, value string) (err error) {
	patchObject := []k8s.PatchPathValue{{
		Op:    "add",
		Path:  "/metadata/annotations/" + annotation,
		Value: value,
	}}
	return r.resource.PatchPath(name, patchObject)
}

func (r *MetaResource) Replace(name string, labels map[string]string, annotations map[string]string) (err error) {
	return r.resource.MergePatch(name, &k8s.MergePatchObject{
		Metadata: &k8s.Metadata{
			Labels:      labels,
			Annotations: annotations,
		},
	})
}
