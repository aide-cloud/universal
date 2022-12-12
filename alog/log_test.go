package alog

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestLog_Debug(t *testing.T) {
	myLog := NewLogger(
		WithOutputType(OutputJsonType),
		WithFileName("log/test.log"),
		WithTimeEncoder(zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000000")),
		WithOutputMode(OutputModeStdoutAndFile),
	)

	myLog.Debug("test", Arg{Key: "key", Value: "value"})
	myLog.Info("test", Arg{Key: "key", Value: "value"})
	myLog.Warn("test", Arg{Key: "key", Value: "value"})
	myLog.Error("test", Arg{Key: "key", Value: "value"})
}
