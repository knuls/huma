package middleware

import (
	"net/http"

	"github.com/knuls/huma/pkg/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObjectID returns a middleware that checks if param is a valid mongo.ObjectID
func ObjectID(param string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if _, err := primitive.ObjectIDFromHex(mux.GetChiURLParam(r, param)); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
