package mux

import "net/http"

type Middleware = func(http.Handler) http.Handler

type Route struct {
	Pattern string
	Handler http.Handler
}

type Multiplexer interface {
	Middlewares(...Middleware)
	Routes(...Route)
	Handler() *http.Handler
}
