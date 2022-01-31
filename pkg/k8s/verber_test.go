package k8s

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"testing"
)

func TestVerberReource(t *testing.T) {
	client := NewResourceVerber(nil, schema.GroupVersionKind{
		Kind:    "application",
		Group:   "apps.gitctl.com",
		Version: "v1",
	})
	client.Namespace("default").Get("demo")
}
