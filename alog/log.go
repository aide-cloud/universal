package alog

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
	"path"
	"runtime"
	"time"
)

type (
	Logger interface {
		// Debug 日志级别为Debug
		Debug(msg string, args ...Arg)
		// Info 日志级别为Info
		Info(msg string, args ...Arg)
		// Warn 日志级别为Warn
		Warn(msg string, args ...Arg)
		// Error 日志级别为Error
		Error(msg string, args ...Arg)
		// Alert 日志级别为Alert
		Alert(hook AlertHook, msg string, args ...Arg)
	}

	Level      string // 日志级别
	OutputType uint8  // 输出类型
	OutputMode uint8  // 输出模式

	Option    func(*Log)                    // 日志配置选项
	AlertHook func(msg string, args ...Arg) // 报警钩子

	Arg struct {
		Key   string // 日志字段名
		Value any    // 日志字段值
	}

	FileLogWriterConfig struct {
		FileName   string // 日志文件名
		MaxSize    int    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups int    // 日志文件最多保存多少个备份
		MaxAge     int    // 文件最多保存多少天
		Compress   bool   // 是否压缩
		LocalTime  bool   // 是否使用本地时间
	}

	Log struct {
		FileLogWriterConfig                     // 文件日志配置
		logger              *zap.Logger         // zap日志对象
		timeEncoder         zapcore.TimeEncoder // 时间编码器
		outputType          OutputType          // 输出类型 1:JSON 2:Console
		outMode             OutputMode          // 日志输出方式 1. file 2. console 3. file+console
	}

	GormLogger struct {
		log      Logger
		LogLevel logger.LogLevel
	}
)

// getArgs build logger args
func getArgs(i []any) []Arg {
	length := len(i)
	var args = make([]Arg, 0, length)
	for j := 0; j < length; j += 2 {
		argTmp := Arg{}
		switch i[j].(type) {
		case string:
			argTmp.Key = i[j].(string)
		default:
			argTmp.Key = "key_" + string(rune(j))
		}
		if j+1 < length {
			args = append(args, Arg{Key: argTmp.Key, Value: i[j+1]})
		} else {
			args = append(args, Arg{Key: argTmp.Key, Value: argTmp.Key})
		}
	}

	return args
}

func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLog := *g
	newLog.LogLevel = level
	return &newLog
}

func (g *GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Info {
		g.log.Info(s, getArgs(i)...)
	}
}

func (g *GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Warn {
		g.log.Warn(s, getArgs(i)...)
	}
}

func (g *GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Error {
		g.log.Error(s, getArgs(i)...)
	}
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.LogLevel > 0 {
		elapsed := time.Since(begin)
		sql, rows := fc()
		if err != nil {
			g.Error(ctx, err.Error(), "rows", rows, "sql", sql, "elapsed", float64(elapsed.Nanoseconds())/1e6)
		} else {
			g.Info(ctx, "", "rows", rows, "sql", sql, "elapsed", float64(elapsed.Nanoseconds())/1e6)
		}
	}
}

func (l *Log) Alert(hook AlertHook, msg string, args ...Arg) {
	go hook(msg, args...)
	l.logger.Fatal(msg, buildLoggerArgs(args, LogLeveLAlert)...)
}

func (l *Log) Debug(msg string, args ...Arg) {
	l.logger.Debug(msg, buildLoggerArgs(args, LogLevelDebug)...)
}

func (l *Log) Info(msg string, args ...Arg) {
	l.logger.Info(msg, buildLoggerArgs(args, LogLevelInfo)...)
}

func (l *Log) Warn(msg string, args ...Arg) {
	l.logger.Warn(msg, buildLoggerArgs(args, LogLevelWarn)...)
}

func (l *Log) Error(msg string, args ...Arg) {
	l.logger.WithOptions(zap.AddStacktrace(zapcore.ErrorLevel)).Error(msg, buildLoggerArgs(args, LogLevelError)...)
}

var _ Logger = (*Log)(nil)
var _ logger.Interface = (*GormLogger)(nil)

// buildLoggerArgs build logger args
func buildLoggerArgs(args []Arg, level Level) []zap.Field {
	callerFields := getCallerInfoForLog()
	fields := make([]zap.Field, 0, len(args)+len(callerFields))
	for _, arg := range args {
		fields = append(fields, anyToZapField(arg.Key, arg.Value))
	}
	if level != LogLevelInfo {
		fields = append(fields, callerFields...)
	}

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
	pc, file, line, ok := runtime.Caller(3) // 回溯3层，拿到写日志的调用方的函数信息
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName) //Base函数返回路径的最后一个元素，只保留函数名
	runtime.Callers(3, []uintptr{pc})

	callerFields = append(
		callerFields,
		zap.String("func", funcName),
		zap.String("file", file),
		zap.Int("line", line),
	)
	return
}
