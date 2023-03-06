package creator

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type creatorIDCtxKey struct{}

func CreatorCtx(key string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), creatorIDCtxKey{}, chi.URLParam(r, key))
			next.ServeHTTP(w, r.Clone(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
