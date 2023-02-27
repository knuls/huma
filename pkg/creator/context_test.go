package creator_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/knuls/huma/pkg/creator"
)

func TestCreatorContext(t *testing.T) {
	mux := chi.NewRouter()
	mux.Use(creator.CreatorCtx)
	mux.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello %s", chi.URLParam(r, "id"))
	})

	req, _ := http.NewRequest(http.MethodGet, "/12345", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	res := rr.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status code is not %d, got %d", http.StatusOK, res.StatusCode)
	}
}
