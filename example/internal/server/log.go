package server

import "github.com/aide-cloud/universal/alog"

var globalLog = alog.NewLogger(
//alog.WithOutputType(alog.OutputJsonType),
)

func GetGlobalLog() alog.Logger {
	return globalLog
}
