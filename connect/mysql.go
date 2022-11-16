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

type (
	// MysqlConfig 数据库配置类型
	MysqlConfig struct {
		User     string `json:"user" yaml:"user"`
		Password string `json:"password" yaml:"password"`
		Host     string `json:"host" yaml:"host"`
		Port     uint   `json:"port" yaml:"port"`
		Db       string `json:"db" yaml:"db"`
		Charset  string `json:"charset" yaml:"charset"`
		Debug    bool   `json:"debug" yaml:"debug"`
		log      logger.Interface
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
		fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		cfg.Db,
		cfg.Charset,
	)

	myLog := cfg.log
	if myLog == nil {
		myLog = alog.NewGormLogger()
	}
	conn, err := gorm.Open(mysql.Open(args), &gorm.Config{Logger: myLog})
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

// WithMysqlLog 设置数据库日志
func WithMysqlLog(log logger.Interface) MysqlConfigOption {
	return func(cfg *MysqlConfig) {
		cfg.log = log
	}
}

// WithMysqlAddr 设置数据库地址
func WithMysqlAddr(host string) MysqlConfigOption {
	return func(cfg *MysqlConfig) {
		cfg.Host = host
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
		cfg.Db = dbName
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
