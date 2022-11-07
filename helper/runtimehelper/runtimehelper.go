package runtimehelper

import (
	"errors"
	"fmt"
	"log"
)

type (
	// RecoverOption is the option for Recover.
	RecoverOption struct {
		log      *log.Logger
		callback func(error)
	}

	RecoverOptionFunc func(*RecoverOption)
)

// NewRecoverOption returns a new RecoverOption.
func NewRecoverOption(options ...RecoverOptionFunc) *RecoverOption {
	rec := &RecoverOption{
		log: log.Default(),
	}

	for _, option := range options {
		option(rec)
	}

	return rec
}

// WithLog sets the log.
func WithLog(log *log.Logger) RecoverOptionFunc {
	return func(rec *RecoverOption) {
		rec.log = log
	}
}

// WithCallback sets the callback.
func WithCallback(callback func(error)) RecoverOptionFunc {
	return func(rec *RecoverOption) {
		rec.callback = callback
	}
}

func Recover(msg string, options *RecoverOption) {
	if err := recover(); err != nil && options != nil {
		recoverErr := errors.New(fmt.Sprintf("%s: %v", msg, err))
		if options.log != nil {
			options.log.Println(recoverErr)
		}

		if options.callback != nil {
			options.callback(recoverErr)
		}
	}
}
