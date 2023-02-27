package logger

import (
	"log"
	"os"
)

type StandardLogger struct {
	log *log.Logger
}

func (l *StandardLogger) Info(message string, args ...interface{}) {
	l.log.Println(message, args)
}

func (l *StandardLogger) Warn(message string, args ...interface{}) {
	l.log.Println(message, args)
}

func (l *StandardLogger) Error(message string, args ...interface{}) {
	l.log.Println(message, args)
}

func (l *StandardLogger) Infof(template string, values ...interface{}) {
	l.log.Printf(template, values...)
}

func (l *StandardLogger) Warnf(template string, values ...interface{}) {
	l.log.Printf(template, values...)
}

func (l *StandardLogger) Errorf(template string, values ...interface{}) {
	l.log.Printf(template, values...)
}

func (l *StandardLogger) GetStdLogger() *log.Logger {
	return l.log
}

func NewStandardLogger() *StandardLogger {
	return &StandardLogger{
		log: log.New(os.Stdout, "", 0),
	}
}
