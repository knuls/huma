package log

import (
	"log"
	"os"
)

type standardLogger struct {
	logger *log.Logger
}

func NewStandardLogger(prefix string) *standardLogger {
	return &standardLogger{
		logger: log.New(os.Stdout, prefix, log.Default().Flags()),
	}
}

func (l *standardLogger) Info(message string, args ...interface{}) {
	l.logger.Println(message, args)
}

func (l *standardLogger) Warn(message string, args ...interface{}) {
	l.logger.Println(message, args)
}

func (l *standardLogger) Debug(message string, args ...interface{}) {
	l.logger.Println(message, args)
}

func (l *standardLogger) Error(message string, args ...interface{}) {
	l.logger.Println(message, args)
}

func (l *standardLogger) GetStdLogger() *log.Logger {
	return l.logger
}
