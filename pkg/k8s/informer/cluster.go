package informer

import (
	"fmt"
	"github.com/gitctl-pro/apps/apis/core/v1"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"time"
)

type ClusterWatcher struct {
	client     rest.Interface
	Workqueue  workqueue.RateLimitingInterface
	Resource   k8s.Resource
	controller cache.Controller
	stopCh     chan struct{}
}

func NewClusterWatcher(config *rest.Config) *ClusterWatcher {
	rateLimit := workqueue.NewItemExponentialFailureRateLimiter(time.Millisecond, 10*time.Second)
	queue := workqueue.NewNamedRateLimitingQueue(rateLimit, "cluster")
	resource := k8s.NewResource(config, &schema.GroupVersionKind{
		Kind: "Cluster", Group: "core.gitctl.com", Version: "v1",
	})

	w := &ClusterWatcher{
		stopCh:    make(chan struct{}),
		Workqueue: queue,
		Resource:  resource,
	}

	informer := k8s.DefaultInformer(resource, &v1.Cluster{}, 0)
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

func (w *ClusterWatcher) Run() {
	defer runtime.HandleCrash()
	go w.controller.Run(w.stopCh)
	if !cache.WaitForCacheSync(w.stopCh, w.controller.HasSynced) {
		runtime.HandleError(fmt.Errorf("Time out waitng for caches to sync"))
		return
	}
	<-w.stopCh
}
