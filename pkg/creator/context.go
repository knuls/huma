package creator

import (
	"context"
	"net/http"

	multiplexer "github.com/knuls/huma/pkg/mux"
)

type creatorIDCtxKey struct{}

func creatorIDCtx(param string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), creatorIDCtxKey{}, multiplexer.GetChiURLParam(r, param))
			next.ServeHTTP(w, r.Clone(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
