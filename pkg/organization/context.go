package organization

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type organizationIDCtxKey struct{}

func OrganizationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), organizationIDCtxKey{}, chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.Clone(ctx))
	})
}
