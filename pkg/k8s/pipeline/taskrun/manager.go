package taskRun

import (
	"context"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"github.com/gitctl-pro/gitctl/pkg/k8s/util/controller"
	"github.com/gitctl-pro/gitctl/pkg/logging"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"time"
)

var (
	log = logging.DefaultLogger.WithField("component", "taskRunManager")
)

type taskRunManager struct {
	config    *rest.Config
	namespace string
	resource  k8s.Resource
	informer  k8s.Informer
}

func NewTaskRunManager(config *rest.Config) k8s.TaskRunManager {
	resource := k8s.NewResource(config, &schema.GroupVersionKind{
		Group:   "tekton.dev",
		Kind:    "TaskRun",
		Version: "v1beta1",
	})
	informer := k8s.NewInformer(resource, &v1beta1.TaskRun{}, 0)
	return &taskRunManager{
		config:   config,
		informer: informer,
		resource: resource,
	}
}

func (m *taskRunManager) Namespace(namespace string) k8s.TaskRunManager {
	m.namespace = namespace
	return m
}

func (m *taskRunManager) Scale(name string, replicas int) error {
	return k8s.NewScale(m.resource).Namespace(m.namespace).ScaleRelicas(name, replicas)
}

func (m *taskRunManager) Watcher(ctx context.Context) {
	log.Info("taskRun watcher")
	go m.informer.Run(ctx.Done())
	go wait.Until(func() {
		workqueue := m.informer.Workqueue()
		resource := m.resource.Resource()
		controller.RunWorker(workqueue, resource, m.handleWatcher)
	}, time.Second, ctx.Done())
	<-ctx.Done()
}

func (m *taskRunManager) handleWatcher(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	obj, exits, err := m.informer.Store().GetByKey(key)
	if errors.IsNotFound(err) || exits {
		log.Info("")
		return nil
	}
	log.Info(namespace, name, obj.(v1beta1.TaskRun))
	return nil
}
