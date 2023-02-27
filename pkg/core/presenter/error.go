package presenter

import (
	"net/http"

	"github.com/go-chi/render"
)

type Error struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// Render implements the chi.Render interface for HTTP payload responses.
func (e *Error) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

// Err returns a render.Renderer generic error HTTP response.
func Err(err error, code int) render.Renderer {
	return &Error{
		Err:        err,
		StatusCode: code,
		Message:    err.Error(),
	}
}

// ErrRender returns a RenderError HTTP response.
func ErrRender(err error) render.Renderer {
	return Err(err, http.StatusUnprocessableEntity)
}

// ErrDecode returns a DecodeError HTTP response.
func ErrDecode(err error) render.Renderer {
	return Err(err, http.StatusBadRequest)
}

// ErrNotFound returns a NotFoundError HTTP response.
func ErrNotFound(err error) render.Renderer {
	return Err(err, http.StatusNotFound)
}

// ErrBadRequest returns a BadRequestError HTTP response.
func ErrBadRequest(err error) render.Renderer {
	return Err(err, http.StatusBadRequest)
}
