package log

import (
	"log"
	"os"
	"sync"
)

type (
	Level      int
	LierLogger struct {
		lock  sync.Mutex
		isDev bool
		log   *log.Logger
	}

	LierLoggerInterface interface {
		Trace(args ...interface{})
		Tracef(format string, args ...interface{})
		Debug(args ...interface{})
		Debugf(format string, args ...interface{})
		Info(args ...interface{})
		Infof(format string, args ...interface{})
		Warn(args ...interface{})
		Warnf(format string, args ...interface{})
		Error(args ...interface{})
		Errorf(format string, args ...interface{})
		Fatal(args ...interface{})
		Fatalf(format string, args ...interface{})
	}
)

// NewLierLogger 创建一个新的日志对象
func NewLierLogger(path string) *LierLogger {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	return &LierLogger{
		log: log.New(file, "", log.LstdFlags),
	}
}

// Trace 记录trace级别的日志
func (l *LierLogger) Trace(args ...interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.isDev {
		log.Println(args...)
		return
	}
	l.log.Println(args...)
}

// Tracef 记录trace级别的日志
func (l *LierLogger) Tracef(format string, args ...interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.isDev {
		log.Printf(format, args...)
		return
	}
	l.log.Printf(format, args...)
}
