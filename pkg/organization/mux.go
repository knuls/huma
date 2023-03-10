package organization

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
	router := chi.NewRouter()
	router.Get("/", m.Find) // GET /organization
	router.Route("/{id}", func(router chi.Router) {
		router.Use(middleware.ValidateObjectID("id"))
		router.Use(OrganizationCtx("id"))
		router.Get("/", m.FindById) // GET /organization/:id
	})
	return router
}

func (m *mux) Find(rw http.ResponseWriter, r *http.Request) {
	result, err := m.svc.Find(r.Context())
	if err != nil {
		render.Render(rw, r, presenter.ErrBadRequest(err))
		return
	}
	orgs := []render.Renderer{}
	for _, org := range result {
		orgs = append(orgs, org)
	}
	render.Status(r, http.StatusOK)
	render.Render(rw, r, &presenter.JSON{"organizations": orgs})
}

func (m *mux) FindById(rw http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(OrganizationIDCtxKey{}).(string)
	org, err := m.svc.FindById(r.Context(), id)
	if err != nil {
		render.Render(rw, r, presenter.ErrBadRequest(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(rw, r, &presenter.JSON{"organization": org})
}
