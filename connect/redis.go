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
		Network            string        `json:"network" yaml:"network"`
		Addr               string        `json:"addr" yaml:"addr"`
		Password           string        `json:"password" yaml:"password"`
		DB                 int           `json:"db" yaml:"db"`
		MaxRetries         int           `json:"max_retries" yaml:"max_retries"`
		MinRetryBackoff    time.Duration `json:"min_retry_backoff" yaml:"min_retry_backoff"`
		MaxRetryBackoff    time.Duration `json:"max_retry_backoff" yaml:"max_retry_backoff"`
		DialTimeout        time.Duration `json:"dial_timeout" yaml:"dial_timeout"`
		ReadTimeout        time.Duration `json:"read_timeout" yaml:"read_timeout"`
		WriteTimeout       time.Duration `json:"write_timeout" yaml:"write_timeout"`
		PoolSize           int           `json:"pool_size" yaml:"pool_size"`
		MinIdleConns       int           `json:"min_idle_conns" yaml:"min_idle_conns"`
		MaxConnAge         time.Duration `json:"max_conn_age" yaml:"max_conn_age"`
		PoolTimeout        time.Duration `json:"pool_timeout" yaml:"pool_timeout"`
		IdleTimeout        time.Duration `json:"idle_timeout" yaml:"idle_timeout"`
		IdleCheckFrequency time.Duration `json:"idle_check_frequency" yaml:"idle_check_frequency"`
		TLSConfig          *tls.Config   `json:"tls_config" yaml:"tls_config"`
		Dialer             func() (net.Conn, error)
		OnConnect          func(*redis.Conn) error
	}

	RedisConfigOption func(*RedisConfig)
)

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
	client := redis.NewClient(&redis.Options{
		Network:            cfg.Network,
		Addr:               cfg.Addr,
		Dialer:             cfg.Dialer,
		OnConnect:          cfg.OnConnect,
		Password:           cfg.Password,
		DB:                 cfg.DB,
		MaxRetries:         cfg.MaxRetries,
		MinRetryBackoff:    cfg.MinRetryBackoff,
		MaxRetryBackoff:    cfg.MaxRetryBackoff,
		DialTimeout:        cfg.DialTimeout,
		ReadTimeout:        cfg.ReadTimeout,
		WriteTimeout:       cfg.WriteTimeout,
		PoolSize:           cfg.PoolSize,
		MinIdleConns:       cfg.MinIdleConns,
		MaxConnAge:         cfg.MaxConnAge,
		PoolTimeout:        cfg.PoolTimeout,
		IdleTimeout:        cfg.IdleTimeout,
		IdleCheckFrequency: cfg.IdleCheckFrequency,
		TLSConfig:          cfg.TLSConfig,
	})

	return client
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
