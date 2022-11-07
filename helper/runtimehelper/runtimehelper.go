package runtimehelper

import (
	"errors"
	"fmt"
	"log"
)

type (
	// RecoverConfig is the option for Recover.
	RecoverConfig struct {
		Log      *log.Logger
		Callback func(error)
	}

	RecoverOption func(*RecoverConfig)
)

// NewRecoverOption returns a new RecoverOption.
func NewRecoverOption(options ...RecoverOption) *RecoverConfig {
	rec := &RecoverConfig{}

	for _, option := range options {
		option(rec)
	}

	if rec.Log == nil {
		WithLog(log.Default())(rec)
	}

	return rec
}

// WithLog sets the log.
func WithLog(log *log.Logger) RecoverOption {
	return func(rec *RecoverConfig) {
		rec.Log = log
	}
}

// WithCallback sets the callback.
func WithCallback(callback func(error)) RecoverOption {
	return func(rec *RecoverConfig) {
		rec.Callback = callback
	}
}

func Recover(msg string, cfg *RecoverConfig) {
	if err := recover(); err != nil && cfg != nil {
		recoverErr := errors.New(fmt.Sprintf("%s: %v", msg, err))
		if cfg.Log != nil {
			cfg.Log.Println(recoverErr)
		}

		if cfg.Callback != nil {
			cfg.Callback(recoverErr)
		}
	}
}
