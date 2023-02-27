package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRealIP(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RemoteAddr == "" {
			t.Fatal("remote ip address not set")
		}
		if r.RemoteAddr != "3.3.3.3" {
			t.Fatalf("remote ip address is incorrect")
		}
	})
	middleware := RealIP(next)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("true-client-ip", "1.1.1.1")
	req.Header.Set("x-real-ip", "2.2.2.2")
	req.Header.Set("x-forwarded-for", "3.3.3.3")
	rr := httptest.NewRecorder()
	middleware.ServeHTTP(rr, req)
}
