package util

import (
	"errors"
	"github.com/gitctl-pro/gitctl/pkg/logging"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	log = logging.DefaultLogger.WithField("component", "clustermanager")
)

func LoadConfig(kubeconfig string) (*rest.Config, error) {
	cfg, err := clientcmd.Load([]byte(kubeconfig))
	if err != nil {
		log.WithError(err).Error("unmarshal: %v", err)
	}
	for context := range cfg.Contexts {
		log.Infof("* %s", context)
		contextCfg, err := clientcmd.NewNonInteractiveClientConfig(*cfg, context, &clientcmd.ConfigOverrides{}, nil).ClientConfig()
		if err != nil {
			log.WithError(err).Error("create %s client: %v", context, err)
		}
		// An arbitrary high number we expect to not exceed. There are various components that need more than the default 5 QPS/10 Burst, e.G.
		// hook for creating ProwJobs and Plank for creating Pods.
		contextCfg.QPS = 100
		contextCfg.Burst = 1000
		for _, v := range cfg.AuthInfos {
			//contextCfg.Username = username
			contextCfg.BearerToken = v.Token
		}

		return contextCfg, nil
	}
	return nil, errors.New("invalid kubeconfig")
}
