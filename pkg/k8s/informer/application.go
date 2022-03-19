package informer

import (
	"fmt"
	"github.com/gitctl-pro/apps/apis/apps/v1"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"time"
)

type ApplicationWatcher struct {
	client    rest.Interface
	Workqueue workqueue.RateLimitingInterface
	informer  cache.Controller
	stopCh    chan struct{}
	resource  k8s.Resource
}

func NewApplicationWatcher(config *rest.Config) *ApplicationWatcher {
	rateLimit := workqueue.NewItemExponentialFailureRateLimiter(time.Millisecond, 10*time.Second)
	queue := workqueue.NewNamedRateLimitingQueue(rateLimit, "application")
	resource := k8s.NewResource(config, &schema.GroupVersionKind{
		Kind:    "Application",
		Group:   "core.gitctl.com",
		Version: "v1",
	})

	w := &ApplicationWatcher{
		stopCh:    make(chan struct{}),
		Workqueue: queue,
		resource:  resource,
	}

	informer := k8s.DefaultInformer(resource, &v1.Application{}, 0)
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

func (w *ApplicationWatcher) Run() {
	defer runtime.HandleCrash()
	go w.informer.Run(w.stopCh)
	if !cache.WaitForCacheSync(w.stopCh, w.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Time out waitng for caches to sync"))
		return
	}
	<-w.stopCh
}
