package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestValidateObjectID(t *testing.T) {
	// target
	id := primitive.NewObjectIDFromTimestamp(time.Now())
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/%s", id.Hex()), nil)
	rr := httptest.NewRecorder()
	mux := chi.NewRouter()
	mux.Use(ValidateObjectID("id"))
	mux.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello %s", chi.URLParam(r, "id"))
	})
	mux.ServeHTTP(rr, req)

	// assert
	res := rr.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status code is not %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
}
