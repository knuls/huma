package mux

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiMux struct {
	mux *chi.Mux
}

func NewChiMux() *chiMux {
	mux := chi.NewRouter()
	return &chiMux{mux: mux}
}

func (m *chiMux) Middlewares(middlewares ...Middleware) {
	m.mux.Use(middlewares...)
}

func (m *chiMux) Routes(routes ...Route) {
	for _, r := range routes {
		m.mux.Mount(r.Pattern, r.Handler)
	}
}

func (m *chiMux) Handler() http.Handler {
	return m.mux
}
