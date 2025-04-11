package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"lowcode/internal/config"
	"lowcode/internal/version"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// show version when start
	logx.Info("version = ", version.Version())
	// show mysql dsn
	logx.Info("mysql dsn = ", c.MYSQL.DSN)

	db, err := gorm.Open(mysql.Open(c.MYSQL.DSN),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	if err != nil {
		panic("Database cannot connected!")
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
