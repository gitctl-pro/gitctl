package pod

import (
	"github.com/gitctl-pro/gitctl/pkg/k8s/logs"
	"io"
	"k8s.io/client-go/rest"
)

func GetLogDetails(cfg *rest.Config, namespace string, name string, container string, logSelector *logs.Selection, usePreviousLogs bool) (*logs.LogDetails, error) {
	podResource := NewPodResource(cfg)
	pod, err := podResource.Namespace(namespace).Name(name).GetPod()

	if err != nil {
		return nil, err
	}

	if len(container) == 0 {
		container = pod.Spec.Containers[0].Name
	}
	fromBegin := true
	if logSelector.LogFilePosition == logs.End {
		fromBegin = false
	}
	readCloser, err := podResource.LogStream(container, fromBegin, usePreviousLogs)
	if err != nil {
		return nil, err
	}
	defer readCloser.Close()
	result, err := io.ReadAll(readCloser)

	if err != nil {
		return nil, err
	}
	details := ConstructLogDetails(name, string(result), container, logSelector)
	return details, nil
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
