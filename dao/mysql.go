package dao

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
	// MysqlDao 数据库操作类型
	MysqlDao struct {
		User     string
		Password string
		Addr     string
		Port     uint
		DBName   string
		Charset  string
		Debug    bool
	}
)

// NewMysqlDao 获取新的数据库连接
func NewMysqlDao(dao *MysqlDao) *gorm.DB {
	if dao == nil {
		panic(errors.New("mysql config is nil"))
	}
	// mysql: 数据库的驱动名
	// 链接数据库 --格式: 用户名:密码@协议(IP:port)/数据库名？xxx&yyy&
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		dao.User,
		dao.Password,
		fmt.Sprintf("%s:%d", dao.Addr, dao.Port),
		dao.DBName,
		dao.Charset,
	)
	conn, err := gorm.Open(mysql.Open(args))
	if err != nil {
		panic(fmt.Sprintf("gorm.Open err: %s", err.Error()))
	}
	if dao.Debug {
		return conn.Debug()
	}
	return conn
}

// NewMysqlDaoSingleton 获取同一个数据库连接
func NewMysqlDaoSingleton(dao *MysqlDao) *gorm.DB {
	if mysqlDB == nil {
		once.Do(func() {
			mysqlDB = NewMysqlDao(dao)
		})
	}
	if dao.Debug {
		return mysqlDB.Debug()
	}
	return mysqlDB
}
