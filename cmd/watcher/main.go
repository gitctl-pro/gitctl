package main

import (
	"flag"
	"github.com/gitctl-pro/gitctl/cmd/dashboard/setup"
	"github.com/gitctl-pro/gitctl/pkg/logging"
	"github.com/spf13/pflag"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/klog/v2"
)

type options struct {
	master     string
	kubeConfig string
	configFile string
	dataDir    string
}

var (
	log = logging.DefaultLogger.WithField("component", "dashboard")
	o   options
)

func InitFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.master, "master", "", "master url")
	fs.StringVar(&o.kubeConfig, "kubeconfig", "", "Path to kubeconfig. Only required if out of cluster")
	fs.StringVar(&o.configFile, "config", "config/route.yml", "")
	fs.StringVar(&o.dataDir, "data", "data/", "")
}

func init() {
	logging.InitLogger()
	klog.InitFlags(nil)
}

func main() {
	InitFlags(pflag.CommandLine)
	pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	config := setup.InitConfig()
	//go setup.SetupMysqlConn(config)
	watcher, _ := setup.SetupK8sWacther(config, o.kubeConfig)
	watcher.EnableDeploymentWatcher()
	watcher.Run()
	//	watcher.EnablePodWatcher()
	//	watcher.EnableEventWatcher()
	//	watcher.EnableReplicasetWatcher()
}
