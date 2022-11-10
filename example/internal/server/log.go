package server

import "github.com/aide-cloud/universal/alog"

var globalLog = alog.NewLogger()

func GetGlobalLog() alog.Logger {
	return globalLog
}
