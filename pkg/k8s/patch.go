package k8s

import (
	"encoding/json"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
)

type ResourcePatch struct {
	resource Resource
	client   rest.Interface
}

type Metadata struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type MergePatchObject struct {
	Metadata *Metadata   `json:"metadata"`
	Spec     interface{} `json:"spec"`
}

type PatchPathValue struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

func (r *resource) MergePatch(name string, patchObject *MergePatchObject) (err error) {
	patchData, _ := json.Marshal(patchObject)
	return r.Patch(name, types.MergePatchType, patchData)
}

func (r *resource) PatchPath(name string, patchObject []PatchPathValue) (err error) {
	patchData, _ := json.Marshal(patchObject)
	return r.Patch(name, types.JSONPatchType, patchData)
}
