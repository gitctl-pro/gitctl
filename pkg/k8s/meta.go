package k8s

import (
	"k8s.io/client-go/rest"
)

type meta struct {
	resource Resource
	client   rest.Interface
}

func NewMeta(resource Resource) MetaResource {
	return &meta{
		resource: resource,
	}
}

func (r *meta) Namespace(namespace string) MetaResource {
	r.resource.Namespace(namespace)
	return r
}

func (r *meta) AddLabel(name string, label, value string) (err error) {
	patchObject := []PatchPathValue{{
		Op:    "add",
		Path:  "/metadata/labels/" + label,
		Value: value,
	}}
	return r.resource.PatchPath(name, patchObject)
}

func (r *meta) RemoveLabel(name string, label string) (err error) {
	patchObject := []PatchPathValue{{
		Op:    "remove",
		Path:  "/metadata/labels/" + label,
		Value: "",
	}}
	return r.resource.PatchPath(name, patchObject)
}

func (r *meta) AddAnnotation(name string, annotation, value string) (err error) {
	patchObject := []PatchPathValue{{
		Op:    "add",
		Path:  "/metadata/annotations/" + annotation,
		Value: value,
	}}
	return r.resource.PatchPath(name, patchObject)
}

func (r *meta) RemoveAnnotation(name string, annotation string) (err error) {
	patchObject := []PatchPathValue{{
		Op:    "remove",
		Path:  "/metadata/annotations/" + annotation,
		Value: "",
	}}
	return r.resource.PatchPath(name, patchObject)
}

func (r *meta) Replace(name string, meta *Metadata) (err error) {
	return r.resource.MergePatch(name, &MergePatchObject{
		Metadata: meta,
	})
}
