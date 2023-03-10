package schema

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/knuls/huma/pkg/core/presenter"
)

type mux struct {
	svc Service
}

func NewMux(svc Service) *mux {
	return &mux{}
}

func (m *mux) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(PopulateOrganizationCtx)
	r.Get("/", m.Find)    // GET /schema
	r.Post("/", m.Create) // POST /schema
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", m.FindById) // GET /schema/:id
		r.Patch("/", m.Update) // PATCH /schema/:id
	})
	return r
}

func (m *mux) Find(rw http.ResponseWriter, r *http.Request) {
	result, err := m.svc.Find(r.Context())
	if err != nil {
		render.Render(rw, r, presenter.ErrBadRequest(err))
		return
	}
	schemas := []render.Renderer{}
	for _, schema := range result {
		schemas = append(schemas, schema)
	}
	render.Status(r, http.StatusOK)
	render.Render(rw, r, &presenter.JSON{"schemas": schemas})
}

func (m *mux) FindById(rw http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.Render(rw, r, &presenter.JSON{"k": "v"})
}

func (m *mux) Create(rw http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.Render(rw, r, &presenter.JSON{"k": "v"})
}

func (m *mux) Update(rw http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.Render(rw, r, &presenter.JSON{"k": "v"})
}
