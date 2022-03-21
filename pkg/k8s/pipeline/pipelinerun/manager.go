package pipelinerun

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
	log = logging.DefaultLogger.WithField("component", "pipelinerunManager")
)

type pipelineRunManager struct {
	config    *rest.Config
	namespace string
	resource  k8s.Resource
	informer  k8s.Informer
}

func NewPipelineRunManager(config *rest.Config) k8s.PipelineRunManager {
	resource := k8s.NewResource(config, &schema.GroupVersionKind{
		Group:   "tekton.dev",
		Kind:    "PipelineRun",
		Version: "v1beta1",
	})
	informer := k8s.NewInformer(resource, &v1beta1.PipelineRun{}, 0)
	return &pipelineRunManager{
		config:   config,
		informer: informer,
		resource: resource,
	}
}

func (m *pipelineRunManager) Namespace(namespace string) k8s.PipelineRunManager {
	m.namespace = namespace
	return m
}

func (m *pipelineRunManager) Scale(name string, replicas int) error {
	return k8s.NewScale(m.resource).Namespace(m.namespace).ScaleRelicas(name, replicas)
}

func (m *pipelineRunManager) Watcher(ctx context.Context) {
	log.Info("pipelinerun watcher")
	go m.informer.Run(ctx.Done())
	go wait.Until(func() {
		workqueue := m.informer.Workqueue()
		resource := m.resource.Resource()
		controller.RunWorker(workqueue, resource, m.handleWatcher)
	}, time.Second, ctx.Done())
	<-ctx.Done()
}

func (m *pipelineRunManager) handleWatcher(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	obj, exits, err := m.informer.Store().GetByKey(key)
	if errors.IsNotFound(err) || exits {
		log.Info("")
		return nil
	}
	log.Info(namespace, name, obj.(v1beta1.PipelineRun))
	return nil
}
