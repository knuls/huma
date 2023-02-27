package creator

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type creatorIDCtxKey struct{}

func CreatorCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), creatorIDCtxKey{}, chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.Clone(ctx))
	})
}
