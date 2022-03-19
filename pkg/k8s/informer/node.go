package informer

import (
	"fmt"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"time"
)

type NodeWatcher struct {
	client     rest.Interface
	Resource   k8s.Resource
	Workqueue  workqueue.RateLimitingInterface
	controller cache.Controller
	StopCh     chan struct{}
}

func NewNodeWatcher(config *rest.Config) *NodeWatcher {
	rateLimit := workqueue.NewItemExponentialFailureRateLimiter(time.Millisecond, 10*time.Second)
	queue := workqueue.NewNamedRateLimitingQueue(rateLimit, "node")
	resource := k8s.NewResource(config, &schema.GroupVersionKind{
		Kind: "node", Version: "v1",
	})
	w := &NodeWatcher{
		StopCh:    make(chan struct{}),
		Workqueue: queue,
		Resource:  resource,
	}

	informer := k8s.DefaultInformer(resource, &v1.Node{}, 0)
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

func (c *NodeWatcher) Run(stopCh chan struct{}) {
	defer runtime.HandleCrash()
	go c.controller.Run(stopCh)
	if !cache.WaitForCacheSync(stopCh, c.controller.HasSynced) {
		runtime.HandleError(fmt.Errorf("Time out waitng for caches to sync"))
		return
	}
	<-stopCh
}
