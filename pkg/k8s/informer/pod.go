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
	client    rest.Interface
	workqueue workqueue.RateLimitingInterface
	informer  cache.Controller
	StopCh    chan struct{}
	resource  k8s.Resource
}

func NewPodWatcher(config *rest.Config) *PodWatcher {
	rateLimit := workqueue.NewItemExponentialFailureRateLimiter(time.Millisecond, 10*time.Second)
	queue := workqueue.NewNamedRateLimitingQueue(rateLimit, "Pods")
	resource := k8s.NewResource(config, &schema.GroupVersionKind{
		Kind:    "pod",
		Version: "v1",
	})

	w := &PodWatcher{
		StopCh:    make(chan struct{}),
		workqueue: queue,
		resource:  resource,
	}

	informerFactory := k8s.NewSharedInformerFactory(resource, 0)
	informer := informerFactory.InformerFor(&v1.Pod{}, k8s.DefaultInformer)

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

func (c *PodWatcher) Run(stopCh chan struct{}) {
	defer runtime.HandleCrash()
	go c.informer.Run(stopCh)
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Time out waitng for caches to sync"))
		return
	}
	<-stopCh
}
