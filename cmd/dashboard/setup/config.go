package setup

import (
	"github.com/gitctl-pro/gitctl/pkg/config"
	"github.com/gitctl-pro/gitctl/pkg/logging"
	"os"
)

var (
	log = logging.DefaultLogger.WithField("component", "console-init")
)

func InitConfig() *config.ConfigResolver {
	log.Info("ParseConfig..")
	yamlFile := "config/config.yml"
	env := os.Getenv("env")
	yamlFile = "config/config-dev.yml"
	if env == "test" {
		yamlFile = "config/config-test.yml"
	} else if env == "pro" {
		yamlFile = "config/config.yml"
	}
	configResolver, err := config.NewResolver(yamlFile)
	if err != nil {
		os.Exit(1)
	}
	return configResolver
}
