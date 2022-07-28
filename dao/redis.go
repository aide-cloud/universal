package dao

// redis 连接池
import (
	"fmt"
	"github.com/go-redis/redis"
)

type (
	// RedisDao redis dao
	RedisDao struct {
		Addr     string
		Port     int
		Password string
		DB       int
	}
)

var (
	redisDao *redis.Client
)

// NewRedisDao 连接池
func NewRedisDao(options *RedisDao) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", options.Addr, options.Port),
		Password: options.Password,
		DB:       options.DB,
	})

	return client
}

// NewRedisDaoSingleton 获取同一个redis连接
func NewRedisDaoSingleton(options *RedisDao) *redis.Client {
	if redisDao == nil {
		once.Do(func() {
			redisDao = NewRedisDao(options)
		})
	}

	return redisDao
}
