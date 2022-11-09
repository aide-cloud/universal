package alog

import "go.uber.org/zap/zapcore"

// WithLevel set log level
func WithLevel(level Level) Option {
	return func(l *Log) {
		l.level = level
	}
}

// WithFileName set log file name
func WithFileName(fileName string) Option {
	return func(l *Log) {
		l.FileName = fileName
	}
}

// WithMaxSize set log file max size
func WithMaxSize(maxSize int) Option {
	return func(l *Log) {
		l.MaxSize = maxSize
	}
}

// WithMaxBackups set log file max backups
func WithMaxBackups(maxBackups int) Option {
	return func(l *Log) {
		l.MaxBackups = maxBackups
	}
}

// WithMaxAge set log file max age
func WithMaxAge(maxAge int) Option {
	return func(l *Log) {
		l.MaxAge = maxAge
	}
}

// WithCompress set log file compress
func WithCompress(compress bool) Option {
	return func(l *Log) {
		l.Compress = compress
	}
}

// WithLocalTime set log file local time
func WithLocalTime(localTime bool) Option {
	return func(l *Log) {
		l.LocalTime = localTime
	}
}

// WithOutputMode set log output mode
func WithOutputMode(mode OutputMode) Option {
	return func(l *Log) {
		l.outMode = mode
	}
}

// WithOutputType set log output type
func WithOutputType(outputType OutputType) Option {
	return func(l *Log) {
		l.outputType = outputType
	}
}

// WithTimeEncoder set log time encoder
func WithTimeEncoder(timeEncoder zapcore.TimeEncoder) Option {
	return func(l *Log) {
		l.timeEncoder = timeEncoder
	}
}
