package pod

import (
	"context"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"io"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

var lineReadLimit int64 = 5000

var byteReadLimit int64 = 500000

type PtyHandler interface {
	io.Reader
	io.Writer
	remotecommand.TerminalSizeQueue
}

type PodResource struct {
	resource  k8s.Resource
	client    rest.Interface
	config    *rest.Config
	namespace string
}

func NewPodResource(cfg *rest.Config) *PodResource {
	resource := k8s.NewResource(cfg, &schema.GroupVersionKind{
		Kind:    "pod",
		Version: "v1",
	})
	return &PodResource{
		client:   resource.Client(),
		resource: resource,
		config:   cfg,
	}

}

func (s *PodResource) Namespace(namespace string) *PodResource {
	s.namespace = namespace
	return s
}

func (s *PodResource) LogStream(name string, container string, fromBegin bool, previous bool) (io.ReadCloser, error) {
	req := s.client.Get().
		Namespace(s.namespace).
		Name(name).
		Resource("pods").
		SubResource("log")

	logOptions := &v1.PodLogOptions{
		Container:  container,
		Previous:   previous,
		Follow:     false,
		Timestamps: true,
	}
	if fromBegin {
		logOptions.LimitBytes = &byteReadLimit
	} else {
		logOptions.TailLines = &lineReadLimit
	}
	req.VersionedParams(logOptions, scheme.ParameterCodec)

	return req.Stream(context.TODO())
}

func (s *PodResource) RemoteCommand(name string, container string, command string, handler PtyHandler) error {
	req := s.client.Post().
		Namespace(s.namespace).
		Name(name).
		Resource("pods").
		SubResource("exec")

	req.VersionedParams(&v1.PodExecOptions{
		Container: container,
		Command:   []string{command},
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(s.config, "POST", req.URL())
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:             handler,
		Stdout:            handler,
		Stderr:            handler,
		Tty:               true,
		TerminalSizeQueue: handler,
	})
	return err
}

func (s *PodResource) Get(name string) (*v1.Pod, error) {
	pod := &v1.Pod{}
	err := s.resource.Namespace(s.namespace).Get(name, pod)
	return pod, err
}
