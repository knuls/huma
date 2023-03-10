package schema

import (
	"context"
	"net/http"
)

type OrganizationNameCtxKey struct{}

func PopulateOrganizationCtx(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// TODO: move this so its re-usable?
		// TODO: fetch org name from:
		// - special org token (for all org related domains)
		// - payload body (for auth only)
		// - url hostname ?
		ctx := context.WithValue(r.Context(), OrganizationNameCtxKey{}, "some-org-name")
		next.ServeHTTP(w, r.Clone(ctx))
	}
	return http.HandlerFunc(fn)
}
