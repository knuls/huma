package log

import (
	"log"
	"log/slog"
	"os"
)

type structuredLogger struct {
	logger *slog.Logger
}

func NewStructuredLogger() *structuredLogger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	return &structuredLogger{
		logger: logger,
	}
}

func (l *structuredLogger) Info(message string, args ...interface{}) {
	l.logger.Info(message, args...)
}

func (l *structuredLogger) Warn(message string, args ...interface{}) {
	l.logger.Warn(message, args...)
}

func (l *structuredLogger) Debug(message string, args ...interface{}) {
	l.logger.Debug(message, args...)
}

func (l *structuredLogger) Error(message string, args ...interface{}) {
	l.logger.Error(message, args...)
}

func (l *structuredLogger) GetStdLogger() *log.Logger {
	return slog.NewLogLogger(l.logger.Handler(), slog.LevelDebug)
}
