package creator

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/knuls/huma/pkg/core/middleware"
	"github.com/knuls/huma/pkg/core/presenter"
)

type mux struct {
	svc Service
}

func NewMux(svc Service) *mux {
	return &mux{
		svc: svc,
	}
}

func (m *mux) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", m.Find) // GET /creator
	r.Route("/{id}", func(r chi.Router) {
		r.Use(middleware.ValidateObjectID("id"))
		r.Use(CreatorCtx("id"))
		r.Get("/", m.FindById) // GET /creator/:id
	})
	return r
}

func (m *mux) Find(rw http.ResponseWriter, r *http.Request) {
	result, err := m.svc.Find(r.Context())
	if err != nil {
		render.Render(rw, r, presenter.ErrBadRequest(err))
		return
	}
	creators := []render.Renderer{}
	for _, creator := range result {
		creators = append(creators, creator)
	}
	render.Status(r, http.StatusOK)
	render.Render(rw, r, &presenter.JSON{"creators": creators})
}

func (m *mux) FindById(rw http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(creatorIDCtxKey{}).(string)
	creator, err := m.svc.FindById(r.Context(), id)
	if err != nil {
		render.Render(rw, r, presenter.ErrBadRequest(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(rw, r, &presenter.JSON{"creator": creator})
}
