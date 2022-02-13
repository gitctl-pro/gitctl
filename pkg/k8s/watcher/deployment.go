package watcher

import (
	"context"
	"fmt"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	runtimeutil "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"time"
)

type DeploymentWatcher struct {
	client    rest.Interface
	workqueue workqueue.RateLimitingInterface
	informer  cache.Controller
	StopCh    chan struct{}
	resource  k8s.Resource
}

func NewFilteredDeploymentInformer(resource k8s.Resource, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				list := &appsv1.DeploymentList{}
				err := resource.List(list, options)
				return list, err
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return resource.Watch(context.TODO(), options)
			},
		},
		&appsv1.Deployment{},
		resyncPeriod,
		indexers,
	)
}

func (w *DeploymentWatcher) defaultInformer(resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDeploymentInformer(w.resource, resyncPeriod, cache.Indexers{})
}

func NewDeploymentWatcher(config *rest.Config) *DeploymentWatcher {

	rateLimit := workqueue.NewItemExponentialFailureRateLimiter(time.Millisecond, 10*time.Second)
	queue := workqueue.NewNamedRateLimitingQueue(rateLimit, "Deployments")
	resource := k8s.NewResource(config, &schema.GroupVersionKind{
		Kind:    "deployment",
		Group:   "apps/v1",
		Version: "v1",
	})

	w := &DeploymentWatcher{
		StopCh:    make(chan struct{}),
		workqueue: queue,
		resource:  resource,
	}

	informerFactory := k8s.NewSharedInformerFactory(resource, 0)
	informer := informerFactory.InformerFor(&appsv1.Deployment{}, w.defaultInformer)

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

func (c *DeploymentWatcher) Run(stopCh chan struct{}) {
	defer runtimeutil.HandleCrash()
	go c.informer.Run(stopCh)
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtimeutil.HandleError(fmt.Errorf("Time out waitng for caches to sync"))
		return
	}
	<-stopCh
}
