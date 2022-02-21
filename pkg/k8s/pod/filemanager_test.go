package pod

import (
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

var (
	kubeConfig = "/Users/zsw/.kube/config"
	cfg, _     = clientcmd.BuildConfigFromFlags("", kubeConfig)
)

func TestLsFile(t *testing.T) {
	file := NewFileManager(cfg, "kube-system", "l7-lb-controller-96f7bf7d8-kljjx", "l7-lb-controller")
	stdout, stderr := file.Ls("/")
	fmt.Println(stdout.String(), stderr.String())
}
