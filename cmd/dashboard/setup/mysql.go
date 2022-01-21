package setup

import (
	"github.com/gitctl-pro/gitctl/pkg/config"
	"github.com/gitctl-pro/gitctl/pkg/gorm"
	"os"
	"time"
)

func SetupMysqlConn(configResolver *config.ConfigResolver) error {
	log.Info("SetupMysqlConn..")
	mysqlOpt, err := configResolver.GetDbConfig("mysql", os.Getenv("env"))
	_, err = gorm.NewMysqlConn(mysqlOpt, 100, 1000, time.Second*1)
	if err != nil {
		log.WithError(err).Error("", err)
		//	os.Exit(1)
	}
	return err
}
