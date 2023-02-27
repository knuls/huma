package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/knuls/huma/pkg/core/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestLogger(t *testing.T) {
	// logger
	logger, err := logger.NewZapLogger()
	if err != nil {
		t.Fatal("logger failed to create")
	}
	defer logger.GetZapLogger().Sync()
	core, capturedLogs := observer.New(zap.InfoLevel)
	sugar := logger.GetZapLogger().Sugar().WithOptions(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return core
	}))
	logger.SetSugar(sugar)

	// target
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("hello world"))
	})
	middleware := Logger(logger)(next)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	middleware.ServeHTTP(rr, req)

	// assert
	expected := map[string]interface{}{
		"proto":  "HTTP/1.1",
		"method": http.MethodGet,
		"path":   "/",
		"status": int64(202),
		"size":   int64(11),
	}
	entry := capturedLogs.All()[0]
	if entry.Level != zap.InfoLevel || entry.Message != "served request" {
		t.Fatal("logger should have written info log with req/res ctx")
	}
	for k, v := range entry.ContextMap() {
		if value, ok := expected[k]; ok && value != v {
			t.Fatalf("logger expected key (%s) value (%v), got (%v)", k, expected[k], v)
		}
	}
	if capturedLogs.Len() != 1 {
		t.Fatal("logger should have captured one log entry")
	}
}
