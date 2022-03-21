package k8s

import (
	"encoding/json"
	"k8s.io/apimachinery/pkg/types"
)

func (r *resource) MergePatch(name string, patchObject *MergePatchObject) (err error) {
	patchData, _ := json.Marshal(patchObject)
	return r.Patch(name, types.MergePatchType, patchData)
}

func (r *resource) PatchPath(name string, patchObject []PatchPathValue) (err error) {
	patchData, _ := json.Marshal(patchObject)
	return r.Patch(name, types.JSONPatchType, patchData)
}
