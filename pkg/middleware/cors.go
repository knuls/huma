package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

func CORS(origins []string, methods []string, headers []string, creds bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		handler := cors.Handler(cors.Options{
			AllowedOrigins:   origins,
			AllowedMethods:   methods,
			AllowedHeaders:   headers,
			AllowCredentials: creds,
		})
		return handler(next)
	}
}
