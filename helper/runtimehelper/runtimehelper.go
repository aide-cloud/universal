package runtimehelper

import (
	"errors"
	"fmt"
	"github.com/aide-cloud/universal/alog"
)

type (
	// RecoverConfig is the option for Recover.
	RecoverConfig struct {
		Log      alog.Logger
		Callback func(error)
	}

	RecoverOption func(*RecoverConfig)
)

// NewRecoverConfig returns a new RecoverOption.
func NewRecoverConfig(options ...RecoverOption) *RecoverConfig {
	rec := &RecoverConfig{}

	for _, option := range options {
		option(rec)
	}

	if rec.Log == nil {
		WithLog(alog.NewLogger())(rec)
	}

	return rec
}

// WithLog sets the log.
func WithLog(log alog.Logger) RecoverOption {
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

func Recover(msg string, cfg ...*RecoverConfig) {
	if err := recover(); err != nil && cfg != nil {
		recoverErr := errors.New(fmt.Sprintf("%s: %v", msg, err))
		if len(cfg) > 0 {
			if cfg[0].Log != nil {
				cfg[0].Log.Error("panic recovered", alog.Arg{Key: "error", Value: recoverErr})
			}

			if cfg[0].Callback != nil {
				cfg[0].Callback(recoverErr)
			}
			return
		}
		alog.NewLogger().Error("panic recovered", alog.Arg{Key: "error", Value: recoverErr})
	}
}
