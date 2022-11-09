package alog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const (
	OutputModeStdout OutputMode = iota + 1
	OutputModeFile
	OutputModeStdoutAndFile
)

const (
	DefaultLogFileName   = "log/a-log.log"
	DefaultLogMaxSize    = 1024
	DefaultLogMaxAge     = 1
	DefaultLogMaxBackups = 5
	DefaultLogCompress   = false
)

const (
	OutputJsonType OutputType = iota + 1
	OutputConsoleType
)

// NewLogger create a new logger
func NewLogger(options ...Option) Logger {
	aLog := Log{}
	// 初始化日志配置项
	for _, option := range options {
		option(&aLog)
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置日志记录中时间的格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if aLog.timeEncoder != nil {
		encoderConfig.EncodeTime = aLog.timeEncoder
	}

	// 日志Encoder 还是JSONEncoder，把日志行格式化成JSON格式的
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	if aLog.outputType == OutputJsonType {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	zapLevel := zapcore.DebugLevel

	cores := make([]zapcore.Core, 0, 2)

	// 设置日志输出位置，支持文件和控制台
	switch aLog.outMode {
	case OutputModeStdout:
		cores = []zapcore.Core{
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapLevel),
		}

	case OutputModeFile:
		cores = []zapcore.Core{
			zapcore.NewCore(encoder, getFileLogWriter(aLog.FileLogWriterConfig), zapLevel),
		}
	case OutputModeStdoutAndFile:
		cores = []zapcore.Core{
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapLevel),
			zapcore.NewCore(encoder, getFileLogWriter(aLog.FileLogWriterConfig), zapLevel),
		}
	default:
		cores = []zapcore.Core{
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapLevel),
		}
	}

	aLog.logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(1))

	return &aLog
}

func getFileLogWriter(config FileLogWriterConfig) (writeSyncer zapcore.WriteSyncer) {
	var cfg = FileLogWriterConfig{
		FileName:   DefaultLogFileName,
		MaxSize:    DefaultLogMaxSize,
		MaxBackups: DefaultLogMaxBackups,
		MaxAge:     DefaultLogMaxAge,
		Compress:   DefaultLogCompress,
	}

	if config.FileName != "" {
		cfg = config
	}

	if config.MaxSize != 0 {
		cfg.MaxSize = config.MaxSize
	}

	if config.MaxBackups != 0 {
		cfg.MaxBackups = config.MaxBackups
	}

	if config.MaxAge != 0 {
		cfg.MaxAge = config.MaxAge
	}

	if config.Compress {
		cfg.Compress = config.Compress
	}

	// 使用 lumberjack 实现 log rotate
	lumberJackLogger := &lumberjack.Logger{
		Filename:   cfg.FileName,
		MaxSize:    cfg.MaxSize,    // 单个文件最大100M
		MaxBackups: cfg.MaxBackups, // 多于 60 个日志文件后，清理较旧的日志
		MaxAge:     cfg.MaxAge,     // 一天一切割
		Compress:   cfg.Compress,   // 是否压缩
		LocalTime:  cfg.LocalTime,  // 使用本地时间
	}

	return zapcore.AddSync(lumberJackLogger)
}
