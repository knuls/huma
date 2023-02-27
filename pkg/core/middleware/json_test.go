package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSON(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content := w.Header().Get("content-type")
		if content != "application/json" {
			t.Fatal("content type is not application/json")
		}
	})
	middleware := JSON(next)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	middleware.ServeHTTP(rr, req)
}
