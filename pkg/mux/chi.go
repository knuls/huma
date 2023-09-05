package mux

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiMux struct {
	mux *chi.Mux
}

func NewChiMux() *chiMux {
	return &chiMux{mux: chi.NewRouter()}
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

func (m *chiMux) Method(method string, pattern string, handlerFn http.HandlerFunc) {
	m.mux.MethodFunc(method, pattern, handlerFn)
}

func (m *chiMux) Fork(pattern string) *chiMux {
	mux := NewChiMux()
	m.Routes(Route{Pattern: pattern, Handler: mux.Handler()})
	return mux
}

func GetChiURLParam(req *http.Request, param string) string {
	return chi.URLParam(req, param)
}
