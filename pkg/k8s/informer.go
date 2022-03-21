package k8s

import (
	"context"
	"github.com/gitctl-pro/gitctl/pkg/k8s/util/controller"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"time"
)

type informer struct {
	client     rest.Interface
	resource   Resource
	workqueue  workqueue.RateLimitingInterface
	controller cache.Controller
	store      cache.Store
}

func NewInformer(resource Resource, object runtime.Object, resyncPeriod time.Duration) *informer {
	rateLimit := workqueue.NewItemExponentialFailureRateLimiter(time.Millisecond, 10*time.Second)
	queue := workqueue.NewNamedRateLimitingQueue(rateLimit, resource.Resource())

	w := &informer{
		workqueue: queue,
		resource:  resource,
	}
	sharedIndexInformer := w.newSharedIndexInformer(object, resyncPeriod)
	sharedIndexInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			controller.Enqueue(obj, queue)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			controller.Enqueue(newObj, queue)
		},
		DeleteFunc: func(obj interface{}) {
			controller.Enqueue(obj, queue)
		},
	})
	w.store = sharedIndexInformer.GetStore()
	return w
}

func (w *informer) newSharedIndexInformer(object runtime.Object, resyncPeriod time.Duration) cache.SharedIndexInformer {
	indexers := cache.Indexers{}
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				obj := &runtime.Unknown{}
				err := w.resource.List(obj, options)
				return obj, err
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return w.resource.Watch(context.TODO(), options)
			},
		},
		object,
		resyncPeriod,
		indexers,
	)
}

func (w *informer) Store() cache.Store {
	return w.store
}

func (w *informer) Workqueue() workqueue.RateLimitingInterface {
	return w.workqueue
}

func (w *informer) Run(stopCh <-chan struct{}) {
	log.Info("informer run:", w.resource)
	go w.controller.Run(stopCh)
	<-stopCh
}
