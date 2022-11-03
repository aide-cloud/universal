package middleware

import (
	"encoding/json"
	"errors"
	"github.com/aide-cloud/universal/cipher"
	"github.com/gin-gonic/gin"
)

const AccessToken = "AccessToken"

type (
	// AesAuthConf is the config for AesAuth
	AesAuthConf[T any] struct {
		HeaderKey    string
		PassCallback PassCallback[T]
		ErrCallback  ErrCallback
		AesCipher    *cipher.AesCipher
	}

	// PassCallback is the callback for pass
	PassCallback[T any] func(args T) bool

	// ErrCallback is the callback for error
	ErrCallback func(ctx *gin.Context, err error)
)

var ErrTokenEmpty = errors.New("token is empty")

// NewAesAuthConf returns a new AesAuthConf
func NewAesAuthConf[T any](headerKey string, passCallback PassCallback[T], errCallback ErrCallback, aesCipher *cipher.AesCipher) AesAuthConf[T] {
	return AesAuthConf[T]{
		HeaderKey:    headerKey,
		PassCallback: passCallback,
		ErrCallback:  errCallback,
		AesCipher:    aesCipher,
	}
}

// AesAuth is a middleware for aes auth
// It will check the token in header
func AesAuth[T any](conf AesAuthConf[T]) gin.HandlerFunc {
	conf.checkConf()
	return func(ctx *gin.Context) {
		var m T

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

		if err = json.Unmarshal([]byte(mStr), &m); err != nil {
			conf.ErrCallback(ctx, err)
			return
		}

		if !conf.PassCallback(m) {
			conf.ErrCallback(ctx, err)
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
		panic("PassCallback is nil")
	}

	if a.ErrCallback == nil {
		panic("ErrCallback is nil")
	}

	if a.AesCipher == nil {
		panic("AesCipher is nil")
	}
}
