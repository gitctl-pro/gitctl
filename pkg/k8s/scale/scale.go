package scale

import (
	"encoding/json"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
)

type scaleResource struct {
	resource k8s.Resource
	client   rest.Interface
}

func NewScaleResource(cfg *rest.Config, gvk *schema.GroupVersionKind) k8s.ScaleResource {
	resource := k8s.NewResource(cfg, gvk)
	return &scaleResource{
		client:   resource.Client(),
		resource: resource,
	}
}

func (r *scaleResource) Namespace(namespace string) k8s.ScaleResource {
	r.resource.Namespace(namespace)
	return r
}

func (r *scaleResource) ScaleRelicas(name string, replicas int) (err error) {
	patchObject := []k8s.PatchPathValue{{
		Op:    "replace",
		Path:  "/spec/replicas",
		Value: replicas,
	}}
	patchData, _ := json.Marshal(patchObject)
	return r.resource.Patch(name, types.JSONPatchType, patchData)
}
