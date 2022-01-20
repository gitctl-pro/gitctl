package gorm

import (
	"errors"
	"fmt"
	"github.com/gitctl-pro/gitctl/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	Db          *gorm.DB
	DbPlat      *gorm.DB
	DbWarehouse *gorm.DB
	//log     = logging.DefaultLogger.WithField("component", "db")
)

type MysqlConn struct {
	Db           *gorm.DB
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  time.Duration
}

func NewMysqlConn(
	mysqlOpt *config.MysqlOpt,
	maxIdleConns int,
	maxOpenConns int,
	maxLifetime time.Duration) (*gorm.DB, error) {
	mysqlConn := &MysqlConn{
		MaxIdleConns: maxIdleConns,
		MaxOpenConns: maxOpenConns,
		MaxLifetime:  maxLifetime,
	}
	db, err := mysqlConn.Setup(mysqlOpt)
	if err != nil {
		//log.Warn(err)
	}
	return db, err
}

func (m *MysqlConn) Setup(opt *config.MysqlOpt) (*gorm.DB, error) {
	if opt == nil {
		return nil, errors.New("cannot get mysqlopt ")
	}
	var dialector gorm.Dialector
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", opt.Username, opt.Password, opt.Host, opt.Port, opt.Database)
	dialector = mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	})
	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	Db, err = gorm.Open(dialector, &gorm.Config{
		//Logger:   logging.NewGormLogger(),
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		//log.Error(err)
	}

	sqlDB, err := Db.DB()

	if err != nil {
		//log.Error("connect db server failed.")
	}
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * 600)
	if err := sqlDB.Ping(); err != nil {
		//log.Error(err)
	}
	return Db, nil
}
