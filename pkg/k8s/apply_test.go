package k8s

import (
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"strings"
	"testing"
)

func TestResourceApply(t *testing.T) {
	body := "{\"kind\":\"Cluster\", \"apiVersion\":\"core.gitctl.com/v1\", \"metadata\": { \"name\": \"test\"}, \"Spec\": {}}"
	reader := strings.NewReader(body)
	d := yaml.NewYAMLOrJSONDecoder(reader, 4096)
	data := &unstructured.Unstructured{}
	if err := d.Decode(data); err != nil {
	}
	gvk := data.GroupVersionKind()
	fmt.Println(gvk.Kind, gvk.Group, gvk.Version)
	resource := NewResource(cfg, &gvk)
	err2 := resource.Create(&runtime.Unknown{Raw: []byte(body), ContentType: runtime.ContentTypeJSON})
	fmt.Println(err2)
}
