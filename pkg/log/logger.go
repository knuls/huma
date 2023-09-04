package log

import "log"

type Logger interface {
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Debug(string, ...interface{})
	Error(string, ...interface{})
	GetStdLogger() *log.Logger
}
