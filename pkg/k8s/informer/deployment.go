package informer

import (
	"fmt"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"time"
)

type DeploymentWatcher struct {
	client     rest.Interface
	Resource   k8s.Resource
	Workqueue  workqueue.RateLimitingInterface
	controller cache.Controller
	stopCh     chan struct{}
}

func NewDeploymentWatcher(config *rest.Config) *DeploymentWatcher {
	rateLimit := workqueue.NewItemExponentialFailureRateLimiter(time.Millisecond, 10*time.Second)
	queue := workqueue.NewNamedRateLimitingQueue(rateLimit, "Deployment")
	resource := k8s.NewResource(config, &schema.GroupVersionKind{
		Kind: "Deployment", Group: "apps", Version: "v1",
	})
	w := &DeploymentWatcher{
		stopCh:    make(chan struct{}),
		Workqueue: queue,
		Resource:  resource,
	}
	informer := k8s.DefaultInformer(resource, &v1.Deployment{}, 0)
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
		},
		DeleteFunc: func(obj interface{}) {
		},
	})
	return w
}

func (w *DeploymentWatcher) Run() {
	defer runtime.HandleCrash()
	go w.controller.Run(w.stopCh)
	if !cache.WaitForCacheSync(w.stopCh, w.controller.HasSynced) {
		runtime.HandleError(fmt.Errorf("Time out waitng for caches to sync"))
		return
	}
	<-w.stopCh
}
