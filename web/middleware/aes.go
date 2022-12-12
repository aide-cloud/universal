package middleware

import (
	"encoding/json"
	"github.com/aide-cloud/universal/basic/assert"
	"github.com/aide-cloud/universal/cipher"
	"github.com/gin-gonic/gin"
)

const (
	AccessToken            = "AccessToken"
	PanicPassCallbackIsNil = "PassCallback is nil"
	PanicErrCallbackIsNil  = "ErrCallback is nil"
)

type (
	// AesAuthConf is the config for AesAuth
	AesAuthConf[T any] struct {
		HeaderKey    string
		PassCallback PassCallback[T]
		ErrCallback  ErrCallback
		AesCipher    *cipher.AesCipher
	}

	AesAuthConfOption[T any] func(conf *AesAuthConf[T])

	// PassCallback is the callback for pass
	// args are: ctx, identity, token, marshal
	// identity is the identity of the user
	// token is the token for the user
	// original is the original data
	PassCallback[T any] func(ctx *gin.Context, identity T, token, original string) bool

	// ErrCallback is the callback for error
	// args are: ctx, err
	//  is the error
	ErrCallback func(ctx *gin.Context, err error)

	BaseError string // BaseError is the base error
)

func (b BaseError) Error() string {
	return string(b)
}

var _ error = BaseError("")

// ErrTokenEmpty is the error for token empty
const ErrTokenEmpty BaseError = "token is empty"

// ErrIdentityType is the error for identity type
const ErrIdentityType BaseError = "'identity' must be a struct or a pointer to a struct"

// ErrTokenInvalid is the error for token invalid
const ErrTokenInvalid BaseError = "token is invalid"

// NewAesAuthConf returns a new AesAuthConf
func NewAesAuthConf[T any](options ...AesAuthConfOption[T]) *AesAuthConf[T] {
	aes := &AesAuthConf[T]{}

	for _, option := range options {
		option(aes)
	}

	aes.checkConf()

	return aes
}

// WithHeaderKey sets the header key
func WithHeaderKey[T any](headKey string) AesAuthConfOption[T] {
	return func(cfg *AesAuthConf[T]) {
		cfg.HeaderKey = headKey
	}
}

// WithPassCallback sets the pass callback
func WithPassCallback[T any](passCallback PassCallback[T]) AesAuthConfOption[T] {
	return func(cfg *AesAuthConf[T]) {
		cfg.PassCallback = passCallback
	}
}

// WithErrCallback sets the error callback
func WithErrCallback[T any](errCallback ErrCallback) AesAuthConfOption[T] {
	return func(cfg *AesAuthConf[T]) {
		cfg.ErrCallback = errCallback
	}
}

// WithAesCipher sets the aes cipher
func WithAesCipher[T any](aesCipher *cipher.AesCipher) AesAuthConfOption[T] {
	return func(cfg *AesAuthConf[T]) {
		cfg.AesCipher = aesCipher
	}
}

// AesAuth is a middleware for aes auth
// It will check the token in header
func AesAuth[T any](conf *AesAuthConf[T]) gin.HandlerFunc {
	conf.checkConf()
	return func(ctx *gin.Context) {
		var m T
		var flag = !assert.IsStruct(m) && !assert.IsStructPtr(m) && !assert.IsArray(m) && !assert.IsSlice(m) && !assert.IsMap(m)

		token := ctx.GetHeader(conf.HeaderKey)
		if token == "" {
			conf.ErrCallback(ctx, ErrTokenEmpty)
			return
		}

		mStr, err := conf.AesCipher.DecryptAesBase64(token)
		if err != nil {
			conf.ErrCallback(ctx, err)
			return
		}

		if !flag {
			if err = json.Unmarshal([]byte(mStr), &m); err != nil {
				conf.ErrCallback(ctx, err)
				return
			}
		}

		if !conf.PassCallback(ctx, m, token, mStr) {
			conf.ErrCallback(ctx, ErrTokenInvalid)
			return
		}
		ctx.Next()
	}
}

func (a AesAuthConf[T]) checkConf() {
	if a.HeaderKey == "" {
		a.HeaderKey = AccessToken
	}

	if a.PassCallback == nil {
		panic(PanicPassCallbackIsNil)
	}

	if a.ErrCallback == nil {
		panic(PanicErrCallbackIsNil)
	}

	if a.AesCipher == nil {
		panic("AesCipher is nil")
	}
}

// GiveGatePass is the pass callback for give gate
func GiveGatePass[T any](identity T, conf *AesAuthConf[T]) gin.HandlerFunc {
	conf.checkConf()
	return func(ctx *gin.Context) {
		if !assert.IsStruct(identity) {
			conf.ErrCallback(ctx, ErrIdentityType)
			return
		}
		//
		marshal, err := json.Marshal(identity)
		if err != nil {
			conf.ErrCallback(ctx, err)
			return
		}

		token, err := conf.AesCipher.EncryptAesBase64(string(marshal))
		if err != nil {
			conf.ErrCallback(ctx, err)
			return
		}

		conf.PassCallback(ctx, identity, token, string(marshal))
	}
}
