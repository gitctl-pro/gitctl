package k8s

import (
	"encoding/json"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
)

type scale struct {
	resource Resource
	client   rest.Interface
}

func NewScale(resource Resource) ScaleResource {
	return &scale{
		resource: resource,
	}
}

func (r *scale) Namespace(namespace string) ScaleResource {
	r.resource.Namespace(namespace)
	return r
}

func (r *scale) ScaleRelicas(name string, replicas int) (err error) {
	patchObject := []PatchPathValue{{
		Op:    "replace",
		Path:  "/spec/replicas",
		Value: replicas,
	}}
	patchData, _ := json.Marshal(patchObject)
	return r.resource.Patch(name, types.JSONPatchType, patchData)
}
