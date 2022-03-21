package k8s

import "k8s.io/client-go/rest"

type Metadata struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type ResourcePatch struct {
	resource Resource
	client   rest.Interface
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
