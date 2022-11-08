package connect

// redis 连接池
import (
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis"
	"net"
	"time"
)

type (
	// RedisConfig redis connect
	RedisConfig struct {
		redis.Options
	}

	RedisConfigOption func(*RedisConfig)
)

var redisCli *redis.Client

// NewRedisConfig 创建redis配置
func NewRedisConfig(options ...RedisConfigOption) *RedisConfig {
	config := &RedisConfig{}
	for _, option := range options {
		option(config)
	}
	return config
}

// NewRedisClient 连接池
func NewRedisClient(cfg *RedisConfig) *redis.Client {
	client := redis.NewClient(&cfg.Options)

	return client
}

// NewRedisDaoSingleton 获取同一个redis连接
func NewRedisDaoSingleton(cfg *RedisConfig) *redis.Client {
	if redisCli == nil {
		once.Do(func() {
			redisCli = NewRedisClient(cfg)
		})
	}

	return redisCli
}

// WithRedisNetwork redis网络类型
func WithRedisNetwork(network string) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.Network = network
	}
}

// WithRedisAddr redis地址
func WithRedisAddr(addr string, port int) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.Addr = fmt.Sprintf("%s:%d", addr, port)
	}
}

// WithRedisDialer redis连接
func WithRedisDialer(dialer func() (net.Conn, error)) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.Dialer = dialer
	}
}

// WithRedisOnConnect redis 连接成功后的回调
func WithRedisOnConnect(onConnect func(*redis.Conn) error) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.OnConnect = onConnect
	}
}

// WithRedisPassword redis密码
func WithRedisPassword(password string) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.Password = password
	}
}

// WithRedisDB redis数据库
func WithRedisDB(db int) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.DB = db
	}
}

// WithRedisMaxRetries redis最大重试次数
func WithRedisMaxRetries(maxRetries int) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.MaxRetries = maxRetries
	}
}

// WithRedisMinRetryBackoff redis最小重试间隔
func WithRedisMinRetryBackoff(minRetryBackoff time.Duration) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.MinRetryBackoff = minRetryBackoff
	}
}

// WithRedisMaxRetryBackoff redis最大重试间隔
func WithRedisMaxRetryBackoff(maxRetryBackoff time.Duration) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.MaxRetryBackoff = maxRetryBackoff
	}
}

// WithRedisDialTimeout redis连接超时时间
func WithRedisDialTimeout(dialTimeout time.Duration) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.DialTimeout = dialTimeout
	}
}

// WithRedisReadTimeout redis读超时时间
func WithRedisReadTimeout(readTimeout time.Duration) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.ReadTimeout = readTimeout
	}
}

// WithRedisWriteTimeout redis写超时时间
func WithRedisWriteTimeout(writeTimeout time.Duration) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.WriteTimeout = writeTimeout
	}
}

// WithRedisPoolSize redis连接池大小
func WithRedisPoolSize(poolSize int) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.PoolSize = poolSize
	}
}

// WithRedisMinIdleConns redis最小空闲连接数
func WithRedisMinIdleConns(minIdleConns int) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.MinIdleConns = minIdleConns
	}
}

// WithRedisMaxConnAge redis最大连接时间
func WithRedisMaxConnAge(maxConnAge time.Duration) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.MaxConnAge = maxConnAge
	}
}

// WithRedisPoolTimeout redis连接池超时时间
func WithRedisPoolTimeout(poolTimeout time.Duration) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.PoolTimeout = poolTimeout
	}
}

// WithRedisIdleTimeout redis空闲超时时间
func WithRedisIdleTimeout(idleTimeout time.Duration) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.IdleTimeout = idleTimeout
	}
}

// WithRedisIdleCheckFrequency redis空闲检查频率
func WithRedisIdleCheckFrequency(idleCheckFrequency time.Duration) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.IdleCheckFrequency = idleCheckFrequency
	}
}

// WithRedisTLSConfig redis tls配置
func WithRedisTLSConfig(tlsConfig *tls.Config) RedisConfigOption {
	return func(cfg *RedisConfig) {
		cfg.TLSConfig = tlsConfig
	}
}
