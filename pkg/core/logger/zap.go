package logger

import (
	"log"

	"go.uber.org/zap"
)

type ZapLogger struct {
	sugar *zap.SugaredLogger
	log   *zap.Logger
}

func (l *ZapLogger) Info(message string, args ...interface{}) {
	l.sugar.Infow(message, args...)
}

func (l *ZapLogger) Warn(message string, args ...interface{}) {
	l.sugar.Warnw(message, args...)
}

func (l *ZapLogger) Error(message string, args ...interface{}) {
	l.sugar.Errorw(message, args...)
}

func (l *ZapLogger) Infof(template string, values ...interface{}) {
	l.sugar.Infof(template, values...)
}

func (l *ZapLogger) Warnf(template string, values ...interface{}) {
	l.sugar.Warnf(template, values...)
}

func (l *ZapLogger) Errorf(template string, values ...interface{}) {
	l.sugar.Errorf(template, values...)
}

func (l *ZapLogger) GetStdLogger() *log.Logger {
	return zap.NewStdLog(l.log)
}

func (l *ZapLogger) GetZapLogger() *zap.Logger {
	return l.log
}

func (l *ZapLogger) SetSugar(sugar *zap.SugaredLogger) {
	l.sugar = sugar
}

func NewZapLogger() (*ZapLogger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &ZapLogger{
		sugar: log.Sugar(),
		log:   log,
	}, nil
}
