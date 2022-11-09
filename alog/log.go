package alog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"runtime"
)

type (
	Logger interface {
		// Debug uses fmt.Sprint to construct and log a message.
		Debug(msg string, args ...Arg)
		// Info uses fmt.Sprint to construct and log a message.
		Info(msg string, args ...Arg)
		// Warn uses fmt.Sprint to construct and log a message.
		Warn(msg string, args ...Arg)
		// Error uses fmt.Sprint to construct and log a message.
		Error(msg string, args ...Arg)
	}

	Level      string
	OutputType uint8
	OutputMode uint8

	Option func(*Log)

	Arg struct {
		Key   string
		Value any
	}

	FileLogWriterConfig struct {
		FileName   string
		MaxSize    int
		MaxBackups int
		MaxAge     int
		Compress   bool
		LocalTime  bool
	}

	Log struct {
		FileLogWriterConfig
		logger      *zap.Logger
		level       Level
		timeEncoder zapcore.TimeEncoder
		outputType  OutputType // 输出类型 1:JSON 2:Console
		outMode     OutputMode // 日志输出方式 1. file 2. console 3. file+console
	}
)

func (l Log) Debug(msg string, args ...Arg) {
	l.logger.Debug(msg, buildLoggerArgs(args)...)
}

func (l Log) Info(msg string, args ...Arg) {
	l.logger.Info(msg, buildLoggerArgs(args)...)
}

func (l Log) Warn(msg string, args ...Arg) {
	l.logger.Warn(msg, buildLoggerArgs(args)...)
}

func (l Log) Error(msg string, args ...Arg) {
	// 打印堆栈信息
	l.logger.Error(msg, buildLoggerArgs(args)...)
}

var _ Logger = (*Log)(nil)

// buildLoggerArgs build logger args
func buildLoggerArgs(args []Arg) []zap.Field {
	callerFields := getCallerInfoForLog()
	fields := make([]zap.Field, 0, len(args)+len(callerFields))
	for _, arg := range args {
		fields = append(fields, anyToZapField(arg.Key, arg.Value))
	}
	fields = append(fields, callerFields...)
	return fields
}

// anyToZapField convert any to zap.Field
func anyToZapField(key string, value any) zap.Field {
	switch v := value.(type) {
	case string:
		return zap.String(key, v)
	case int:
		return zap.Int(key, v)
	case int64:
		return zap.Int64(key, v)
	case uint:
		return zap.Uint(key, v)
	case uint64:
		return zap.Uint64(key, v)
	case float32:
		return zap.Float32(key, v)
	case float64:
		return zap.Float64(key, v)
	case bool:
		return zap.Bool(key, v)
	case error:
		return zap.Error(v)
	case zap.Field:
		return v
	default:
		return zap.Any(key, v)
	}
}

func getCallerInfoForLog() (callerFields []zap.Field) {

	pc, file, line, ok := runtime.Caller(2) // 回溯两层，拿到写日志的调用方的函数信息
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName) //Base函数返回路径的最后一个元素，只保留函数名

	callerFields = append(callerFields, zap.String("func", funcName), zap.String("file", file), zap.Int("line", line))
	return
}
