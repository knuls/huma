package middleware

import (
	"net/http"
	"time"

	"github.com/knuls/huma/pkg/log"
)

// Logger is a func in which returns a middleware such that req & res entries are logged.
func Logger(logger log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			wrw := NewWrapResponseWriter(w, r.ProtoMajor)
			defer func() {
				logger.Info("served request",
					"proto", r.Proto,
					"method", r.Method,
					"path", r.URL.Path,
					"took", time.Since(now),
					"status", wrw.Status(),
					"size", wrw.BytesWritten(),
					"reqId", GetReqID(r.Context()),
				)
			}()
			next.ServeHTTP(wrw, r)
		}
		return http.HandlerFunc(fn)
	}
}
