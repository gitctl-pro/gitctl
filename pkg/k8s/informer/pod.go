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

type PodWatcher struct {
	client     rest.Interface
	Resource   k8s.Resource
	Workqueue  workqueue.RateLimitingInterface
	controller cache.Controller
	stopCh     chan struct{}
}

func NewPodWatcher(config *rest.Config) *PodWatcher {
	rateLimit := workqueue.NewItemExponentialFailureRateLimiter(time.Millisecond, 10*time.Second)
	queue := workqueue.NewNamedRateLimitingQueue(rateLimit, "pod")
	resource := k8s.NewResource(config, &schema.GroupVersionKind{
		Kind: "pod", Version: "v1",
	})

	w := &PodWatcher{
		stopCh:    make(chan struct{}),
		Workqueue: queue,
		Resource:  resource,
	}
	informer := k8s.DefaultInformer(resource, &v1.Pod{}, 0)
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

func (w *PodWatcher) Run(stopCh chan struct{}) {
	defer runtime.HandleCrash()
	go w.controller.Run(stopCh)
	if !cache.WaitForCacheSync(stopCh, w.controller.HasSynced) {
		runtime.HandleError(fmt.Errorf("Time out waitng for caches to sync"))
		return
	}
	<-w.stopCh
}
