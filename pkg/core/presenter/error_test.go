package presenter

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestError(t *testing.T) {
	e := &Error{}
	req := &http.Request{}
	rr := httptest.NewRecorder()
	err := e.Render(rr, req)
	if err != nil {
		t.Error(err)
	}
}

func TestErr(t *testing.T) {
	res := Err(errors.New("no impl"), http.StatusAccepted)
	req := &http.Request{}
	rr := httptest.NewRecorder()
	err := res.Render(rr, req)
	if err != nil {
		t.Error(err)
	}
}

func TestErrRender(t *testing.T) {
	res := ErrRender(errors.New("render"))
	req := &http.Request{}
	rr := httptest.NewRecorder()
	err := res.Render(rr, req)
	if err != nil {
		t.Error(err)
	}
}

func TestErrDecode(t *testing.T) {
	res := ErrDecode(errors.New("decode"))
	req := &http.Request{}
	rr := httptest.NewRecorder()
	err := res.Render(rr, req)
	if err != nil {
		t.Error(err)
	}
}

func TestErrNotFound(t *testing.T) {
	res := ErrNotFound(errors.New("not found"))
	req := &http.Request{}
	rr := httptest.NewRecorder()
	err := res.Render(rr, req)
	if err != nil {
		t.Error(err)
	}
}

func TestErrBadRequest(t *testing.T) {
	res := ErrBadRequest(errors.New("boom"))
	req := &http.Request{}
	rr := httptest.NewRecorder()
	err := res.Render(rr, req)
	if err != nil {
		t.Error(err)
	}
}
