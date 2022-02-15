package k8s

import (
	"context"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"time"
)

func DefaultInformer(resource Resource, object runtime.Object, resyncPeriod time.Duration) cache.SharedIndexInformer {
	indexers := cache.Indexers{}
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				obj := &runtime.Unknown{}
				err := resource.List(obj, options)
				return obj, err
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return resource.Watch(context.TODO(), options)
			},
		},
		object,
		resyncPeriod,
		indexers,
	)
}
