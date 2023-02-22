package middleware

import (
	"errors"
	"github.com/aide-cloud/universal/cipher"
	"github.com/aide-cloud/universal/gin/response"
	"github.com/gin-gonic/gin"
)

const (
	AccessToken = "AccessToken"
)

type (
	// AesAuthConfig 鉴权中间件配置
	AesAuthConfig struct {
		cipherInstance cipher.AesInterface
		resp           response.Interface
		checkToken     func(token []byte) bool
	}

	Option func(*AesAuthConfig)
)

var (
	// ErrTokenEmpty token 为空
	ErrTokenEmpty = errors.New("token is empty")
	// ErrTokenInvalid token 无效
	ErrTokenInvalid = errors.New("token is invalid")
)

// NewAesAuthConfig 鉴权中间件配置
func NewAesAuthConfig(cipherInstance cipher.AesInterface, opts ...Option) *AesAuthConfig {
	config := &AesAuthConfig{
		cipherInstance: cipherInstance,
		resp:           response.NewDefault(),
		checkToken: func(token []byte) bool {
			return true
		},
	}

	for _, opt := range opts {
		opt(config)
	}

	return config
}

// AesAuth 鉴权中间件
func AesAuth(config *AesAuthConfig) func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader(AccessToken)
		if token == "" {
			config.resp.Failed(c, "", ErrTokenEmpty)
			return
		}

		// 解密
		decrypt, err := config.cipherInstance.DecryptBase64(token)
		if err != nil {
			config.resp.Failed(c, ErrTokenInvalid.Error(), err)
			return
		}

		// 校验
		if !config.checkToken(decrypt) {
			config.resp.Failed(c, "", ErrTokenInvalid)
			return
		}

		c.Set(AccessToken, decrypt)
		c.Next()
	}
}

// WithResp 设置响应
func WithResp(resp response.Interface) Option {
	return func(config *AesAuthConfig) {
		config.resp = resp
	}
}

// WithCheckToken 设置校验 token
func WithCheckToken(checkToken func(token []byte) bool) Option {
	return func(config *AesAuthConfig) {
		config.checkToken = checkToken
	}
}
