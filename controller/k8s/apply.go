package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/pkg/controller"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"strings"
)

func (c *K8sController) Apply(ctx *gin.Context) {
	cluster := ctx.Query("cluster")
	namespace := ctx.Query("namespace")
	cfg, err := c.clusterManager.Get(cluster)
	contentType := ctx.ContentType()

	body, _ := ioutil.ReadAll(ctx.Request.Body)
	reader := strings.NewReader(string(body))
	d := yaml.NewYAMLOrJSONDecoder(reader, 4096)
	data := &unstructured.Unstructured{}
	if err := d.Decode(data); err != nil {
	}

	gvk := data.GroupVersionKind()
	resource := k8s.NewResource(cfg, &gvk)
	if len(namespace) > 0 {
		resource.Namespace(namespace)
	}
	err = resource.Get(data.GetName(), &runtime.Unknown{})
	if err != nil {
		err = resource.Create(&runtime.Unknown{Raw: body, ContentType: contentType})
	} else {
		err = resource.Put(data.GetName(), &runtime.Unknown{Raw: body, ContentType: contentType})
	}

	ctx.JSON(200, &controller.Response{
		Err:  err,
		Msg:  "",
		Data: nil,
	})
	return
}
