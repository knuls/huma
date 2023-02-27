package middleware

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/knuls/huma/pkg/core/presenter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ValidateObjectID returns a middleware that checks if the `key` param is a valid mongo.ObjectID
func ValidateObjectID(key string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if !primitive.IsValidObjectID(chi.URLParam(r, key)) {
				err := errors.New("invalid object id param")
				render.Render(w, r, presenter.Err(err, http.StatusBadRequest))
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
