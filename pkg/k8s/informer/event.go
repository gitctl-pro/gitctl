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

type EventWatcher struct {
	client    rest.Interface
	Resource  k8s.Resource
	Workqueue workqueue.RateLimitingInterface
	contoller cache.Controller
	stopCh    chan struct{}
}

func NewEventWatcher(config *rest.Config) *EventWatcher {
	rateLimit := workqueue.NewItemExponentialFailureRateLimiter(time.Millisecond, 10*time.Second)
	queue := workqueue.NewNamedRateLimitingQueue(rateLimit, "event")
	resource := k8s.NewResource(config, &schema.GroupVersionKind{
		Kind: "Event", Version: "v1",
	})

	w := &EventWatcher{
		stopCh:    make(chan struct{}),
		Workqueue: queue,
		Resource:  resource,
	}

	informer := k8s.DefaultInformer(resource, &v1.Event{}, 0)
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

func (w *EventWatcher) Run(stopCh chan struct{}) {
	defer runtime.HandleCrash()
	go w.contoller.Run(stopCh)
	if !cache.WaitForCacheSync(stopCh, w.contoller.HasSynced) {
		runtime.HandleError(fmt.Errorf("Time out waitng for caches to sync"))
		return
	}
	<-stopCh
}
