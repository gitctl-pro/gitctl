package pod

import (
	"bytes"
	"fmt"
	"k8s.io/client-go/rest"
)

type FileManager struct {
	resource  *PodResource
	container string
}

func NewFileManager(cfg *rest.Config, namespace, name, container string) *FileManager {
	resource := NewPodResource(cfg).
		Namespace(namespace).
		Name(name)

	return &FileManager{
		resource:  resource,
		container: container,
	}
}

func (f *FileManager) Ls(path string) (bytes.Buffer, bytes.Buffer) {
	command := "ls -l " + path
	stdout, stderr, err := f.resource.ExecStream(f.container, command, 1, nil)

	if err != nil {
		fmt.Println(err)
	}
	return stdout, stderr
}

func (f *FileManager) Exist(filename string) bool {
	return false
}

func (f *FileManager) Size(filename string) {

}

func (f *FileManager) Tar(filename string) {

}

func (f *FileManager) Upload(filename string) {

}

func (f *FileManager) Download(filename string) {

}

func (f *FileManager) Sync(filename string) {

}
