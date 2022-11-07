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

	MysqlConfigOption func(*MysqlConfig)
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

// NewMysqlConfig 创建数据库配置
func NewMysqlConfig(opts ...MysqlConfigOption) *MysqlConfig {
	cfg := &MysqlConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// WithMysqlUser 设置数据库用户名
func WithMysqlUser(user string) MysqlConfigOption {
	return func(cfg *MysqlConfig) {
		cfg.User = user
	}
}

// WithMysqlPassword 设置数据库密码
func WithMysqlPassword(password string) MysqlConfigOption {
	return func(cfg *MysqlConfig) {
		cfg.Password = password
	}
}

// WithMysqlAddr 设置数据库地址
func WithMysqlAddr(addr string) MysqlConfigOption {
	return func(cfg *MysqlConfig) {
		cfg.Addr = addr
	}
}

// WithMysqlPort 设置数据库端口
func WithMysqlPort(port uint) MysqlConfigOption {
	return func(cfg *MysqlConfig) {
		cfg.Port = port
	}
}

// WithMysqlDBName 设置数据库名
func WithMysqlDBName(dbName string) MysqlConfigOption {
	return func(cfg *MysqlConfig) {
		cfg.DBName = dbName
	}
}

// WithMysqlCharset 设置数据库字符集
func WithMysqlCharset(charset string) MysqlConfigOption {
	return func(cfg *MysqlConfig) {
		cfg.Charset = charset
	}
}

// WithMysqlDebug 设置数据库调试模式
func WithMysqlDebug(debug bool) MysqlConfigOption {
	return func(cfg *MysqlConfig) {
		cfg.Debug = debug
	}
}
