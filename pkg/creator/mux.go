package creator

import (
	"fmt"
	"net/http"

	"github.com/knuls/huma/pkg/middleware"
	multiplexer "github.com/knuls/huma/pkg/mux"
)

type mux struct {
}

func NewMux() *mux {
	return &mux{}
}

func (m *mux) Routes() http.Handler {
	router := multiplexer.NewChiMux()
	router.Method(http.MethodGet, "/", m.Find)    // GET /creator
	router.Method(http.MethodPost, "/", m.Create) // POST /creator

	param := "id"
	subRouter := router.Fork(fmt.Sprintf("/{%s}", param))
	subRouter.Middlewares(
		middleware.ObjectID(param),
		creatorIDCtx(param),
	)
	subRouter.Method(http.MethodGet, "/", m.FindById)   // GET /creator/:id
	subRouter.Method(http.MethodPatch, "/", m.FindById) // PATCH /creator/:id

	return router.Handler()
}

func (m *mux) Find(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("find"))
}

func (m *mux) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create"))
}

func (m *mux) FindById(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(creatorIDCtxKey{}).(string)
	w.Write([]byte(id))
}

func (m *mux) Update(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(creatorIDCtxKey{}).(string)
	w.Write([]byte(id))
}
