package middleware

import "github.com/go-chi/chi/v5/middleware"

// Recoverer is a middleware that recovers from panics, logs the panic (and a
// backtrace), and returns a HTTP 500 (Internal Server Error) status if
// possible. Recoverer prints a request ID if one is provided.
//
// This is a wrapper around go-chi/middleware/Recoverer.
var Recoverer = middleware.Recoverer
