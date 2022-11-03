package connect

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	mysqlDB *gorm.DB
	once    sync.Once
)

type (
	// MysqlConfig 数据库配置类型
	MysqlConfig struct {
		User     string
		Password string
		Addr     string
		Port     uint
		DBName   string
		Charset  string
		Debug    bool
	}
)

// GetMysqlConnect 获取新的数据库连接
func GetMysqlConnect(cfg *MysqlConfig) *gorm.DB {
	if cfg == nil {
		panic(errors.New("mysql config is nil"))
	}
	// mysql: 数据库的驱动名
	// 链接数据库 --格式: 用户名:密码@协议(IP:port)/数据库名？xxx&yyy&
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port),
		cfg.DBName,
		cfg.Charset,
	)
	conn, err := gorm.Open(mysql.Open(args))
	if err != nil {
		panic(fmt.Sprintf("gorm.Open err: %s", err.Error()))
	}
	if cfg.Debug {
		return conn.Debug()
	}
	return conn
}

// GetMysqlConnectSingle 获取同一个数据库连接
func GetMysqlConnectSingle(cfg *MysqlConfig) *gorm.DB {
	if mysqlDB == nil {
		once.Do(func() {
			mysqlDB = GetMysqlConnect(cfg)
		})
	}
	if cfg.Debug {
		return mysqlDB.Debug()
	}
	return mysqlDB
}
