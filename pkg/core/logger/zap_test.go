package logger

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestZapLogger(t *testing.T) {
	logger, err := NewZapLogger()
	if err != nil {
		t.Fatal("logger failed to create")
	}
	defer logger.GetZapLogger().Sync()

	// replace logger zap/core with observed zap/core to capture written logs
	core, capturedLogs := observer.New(zap.InfoLevel)
	sugar := logger.log.Sugar().WithOptions(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return core
	}))
	logger.SetSugar(sugar)

	// write logs
	logger.Info("some log line", "key", "value")
	logger.Warn("some log line", "key", "value")
	logger.Error("some log line", "key", "value")
	logger.Infof("hello %s", "world")
	logger.Warnf("hello %s", "world")
	logger.Errorf("hello %s", "world")

	// assert
	logs := capturedLogs.All()
	entry := logs[0]
	if entry.Level != zap.InfoLevel || entry.Message != "some log line" || entry.ContextMap()["key"] != "value" {
		t.Fatal("logger should have written info log with message and key/value")
	}
	if capturedLogs.Len() != 6 {
		t.Fatal("logger should have captured two log entries")
	}
	if !isZapLogger(logger.GetZapLogger()) {
		t.Fatal("logger should be of type zap/logger")
	}
	if !isStandardLogger(logger.GetStdLogger()) {
		t.Fatal("logger should be of type log/logger")
	}
}

func isZapLogger(o interface{}) bool {
	switch o.(type) {
	case *zap.Logger:
		return true
	default:
		return false
	}
}
