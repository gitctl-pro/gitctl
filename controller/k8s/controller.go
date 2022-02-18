package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/gitctl-pro/gitctl/controller/k8s/apps"
	"github.com/gitctl-pro/gitctl/controller/k8s/autoscaling"
	"github.com/gitctl-pro/gitctl/controller/k8s/batch"
	"github.com/gitctl-pro/gitctl/controller/k8s/core"
	"github.com/gitctl-pro/gitctl/controller/k8s/extension"
	"github.com/gitctl-pro/gitctl/controller/k8s/networking"
	"github.com/gitctl-pro/gitctl/controller/k8s/rabc"
	"github.com/gitctl-pro/gitctl/controller/k8s/storage"
	"github.com/gitctl-pro/gitctl/pkg/controller"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/rest"
	"strings"
)

type K8sController struct {
	clusterManager k8s.ClusterManager
	Deployment     DeploymentInterface
	StatefulSet    StatefulSetInterface
	DaemonSet      DaemonsetInterface
	ReplicaSet     ReplicasetInterface
	Node           NodeInterface
	Namespace      NamespaceInterface
	Pod            PodInterface
	ConfigMap      ConfigmapInterface
	Service        ServiceInterface
	Ingress        IngressInterface
	IngressClass   IngressClassInterface
	NetworkPolicy  NetworkPolicyInterface
	Job            JobInterface
	CronJob        CronjobInterface
	Secret         SecretInterface
	HPA            HPAInterface
	PV             PVInterface
	PVC            PVCInterface
	Event          EventInterface
	Scale          ScaleInterface
	ServiceAccount ServiceAccountInterface
	ClusterRole    ClusterRoleInterface
	Role           RoleInterface
	Crd            CrdInterface
	StorageClass   StorageClassInterface
}

func NewController(cfg *rest.Config, clusterManager k8s.ClusterManager) *K8sController {
	return &K8sController{
		Deployment:     apps.NewDeployment(clusterManager),
		ReplicaSet:     apps.NewReplicaset(clusterManager),
		DaemonSet:      apps.NewDaemonset(clusterManager),
		StatefulSet:    apps.NewStatefulSet(clusterManager),
		Namespace:      core.NewNamespace(clusterManager),
		Node:           core.NewNode(clusterManager),
		Pod:            core.NewPod(clusterManager),
		ConfigMap:      core.NewConfigmap(clusterManager),
		Service:        core.NewService(clusterManager),
		Ingress:        networking.NewIngress(clusterManager),
		IngressClass:   networking.NewIngressClass(clusterManager),
		NetworkPolicy:  networking.NewNetworkPolicy(clusterManager),
		HPA:            autoscaling.NewHPA(clusterManager),
		Job:            batch.NewJob(clusterManager),
		CronJob:        batch.NewCronjob(clusterManager),
		Secret:         core.NewSecret(clusterManager),
		PVC:            core.NewPVC(clusterManager),
		PV:             core.NewPV(clusterManager),
		Event:          core.NewEvent(clusterManager),
		ServiceAccount: core.NewServiceAccount(clusterManager),
		ClusterRole:    rabc.NewClusterRole(clusterManager),
		Role:           rabc.NewRole(clusterManager),
		Crd:            extension.NewCrd(clusterManager),
		StorageClass:   storage.NewStorageClass(clusterManager),
	}
}

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

func (c *K8sController) Patch(ctx *gin.Context) {

}
