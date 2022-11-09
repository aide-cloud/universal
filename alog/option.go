package alog

import "go.uber.org/zap/zapcore"

// WithFileName 设置日志文件名，不包含路径，只有在输出方式为文件或者控制台+文件时有效
func WithFileName(fileName string) Option {
	return func(l *Log) {
		l.FileName = fileName
	}
}

// WithMaxSize 设置日志文件最大尺寸，单位MB
func WithMaxSize(maxSize int) Option {
	return func(l *Log) {
		l.MaxSize = maxSize
	}
}

// WithMaxBackups 设置日志文件最大备份数
func WithMaxBackups(maxBackups int) Option {
	return func(l *Log) {
		l.MaxBackups = maxBackups
	}
}

// WithMaxAge 设置日志文件最大保存天数
func WithMaxAge(maxAge int) Option {
	return func(l *Log) {
		l.MaxAge = maxAge
	}
}

// WithCompress 设置日志文件是否压缩
func WithCompress(compress bool) Option {
	return func(l *Log) {
		l.Compress = compress
	}
}

// WithLocalTime 设置日志文件是否使用本地时间
func WithLocalTime(localTime bool) Option {
	return func(l *Log) {
		l.LocalTime = localTime
	}
}

// WithOutputMode 设置日志输出方式，控制台、文件、控制台+文件
func WithOutputMode(mode OutputMode) Option {
	return func(l *Log) {
		l.outMode = mode
	}
}

// WithOutputType 设置日志输出类型，JSON或者Console
func WithOutputType(outputType OutputType) Option {
	return func(l *Log) {
		l.outputType = outputType
	}
}

// WithTimeEncoder 设置日志时间格式，支持自定义格式
func WithTimeEncoder(timeEncoder zapcore.TimeEncoder) Option {
	return func(l *Log) {
		l.timeEncoder = timeEncoder
	}
}
