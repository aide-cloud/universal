package connect

import (
	"errors"
	"fmt"
	"github.com/aide-cloud/universal/alog"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	mysqlDB *gorm.DB
	once    sync.Once
)

// GetMysqlConnect 获取新的数据库连接
func GetMysqlConnect(dsn string, log ...logger.Interface) *gorm.DB {
	if dsn == "" {
		panic(errors.New("mysql dsn is \"\""))
	}

	var myLog logger.Interface
	if len(log) > 0 {
		myLog = log[0]
	} else {
		myLog = alog.NewGormLogger()
	}

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: myLog})
	if err != nil {
		panic(fmt.Sprintf("gorm.Open err: %s", err.Error()))
	}
	return conn
}

// GetMysqlConnectSingle 获取同一个数据库连接
func GetMysqlConnectSingle(dsn string, log ...logger.Interface) *gorm.DB {
	if mysqlDB == nil {
		once.Do(func() {
			mysqlDB = GetMysqlConnect(dsn, log...)
		})
	}
	return mysqlDB
}
