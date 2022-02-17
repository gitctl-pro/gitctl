package stream

import (
	"context"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"github.com/gitctl-pro/gitctl/pkg/k8s/logs"
	"io"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// maximum number of lines loaded from the apiserver
var lineReadLimit int64 = 5000

// maximum number of bytes loaded from the apiserver
var byteReadLimit int64 = 500000

// PodContainerList is a list of containers of a pod.
type PodContainerList struct {
	Containers []string `json:"containers"`
}

type LogStream struct {
	resource k8s.Resource
	client   rest.Interface
}

func NewLogStream(cfg *rest.Config) *LogStream {
	resource := k8s.NewResource(cfg, &schema.GroupVersionKind{
		Kind:    "pod",
		Version: "v1",
	})
	return &LogStream{
		client:   resource.Client(),
		resource: resource,
	}
}

func (l *LogStream) LogStream(namespace, name string, logOptions *v1.PodLogOptions) (io.ReadCloser, error) {
	return l.client.Get().
		Namespace(namespace).
		Name(name).
		Resource("pods").
		SubResource("log").
		VersionedParams(logOptions, scheme.ParameterCodec).Stream(context.TODO())
}

func (l *LogStream) GetLogDetails(namespace, podID string, container string,
	logSelector *logs.Selection, usePreviousLogs bool) (*logs.LogDetails, error) {
	pod := &v1.Pod{}
	err := l.resource.Namespace(namespace).Get(podID, pod)
	if err != nil {
		return nil, err
	}

	if len(container) == 0 {
		container = pod.Spec.Containers[0].Name
	}

	logOptions := mapToLogOptions(container, logSelector, usePreviousLogs)
	readCloser, err := l.LogStream(namespace, podID, logOptions)
	if err != nil {
		return nil, err
	}
	defer readCloser.Close()
	result, err := io.ReadAll(readCloser)

	if err != nil {
		return nil, err
	}
	details := ConstructLogDetails(podID, string(result), container, logSelector)
	return details, nil
}

// Maps the log selection to the corresponding api object
// Read limits are set to avoid out of memory issues
func mapToLogOptions(container string, logSelector *logs.Selection, previous bool) *v1.PodLogOptions {
	logOptions := &v1.PodLogOptions{
		Container:  container,
		Follow:     false,
		Previous:   previous,
		Timestamps: true,
	}

	if logSelector.LogFilePosition == logs.Beginning {
		logOptions.LimitBytes = &byteReadLimit
	} else {
		logOptions.TailLines = &lineReadLimit
	}

	return logOptions
}

// ConstructLogDetails creates a new log details structure for given parameters.
func ConstructLogDetails(podID string, rawLogs string, container string, logSelector *logs.Selection) *logs.LogDetails {
	parsedLines := logs.ToLogLines(rawLogs)
	logLines, fromDate, toDate, logSelection, lastPage := parsedLines.SelectLogs(logSelector)

	readLimitReached := isReadLimitReached(int64(len(rawLogs)), int64(len(parsedLines)), logSelector.LogFilePosition)
	truncated := readLimitReached && lastPage

	info := logs.LogInfo{
		PodName:       podID,
		ContainerName: container,
		FromDate:      fromDate,
		ToDate:        toDate,
		Truncated:     truncated,
	}
	return &logs.LogDetails{
		Info:      info,
		Selection: logSelection,
		LogLines:  logLines,
	}
}

// Checks if the amount of log file returned from the apiserver is equal to the read limits
func isReadLimitReached(bytesLoaded int64, linesLoaded int64, logFilePosition string) bool {
	return (logFilePosition == logs.Beginning && bytesLoaded >= byteReadLimit) ||
		(logFilePosition == logs.End && linesLoaded >= lineReadLimit)
}
