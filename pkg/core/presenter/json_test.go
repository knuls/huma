package presenter

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSON(t *testing.T) {
	j := &JSON{}
	req := &http.Request{}
	rr := httptest.NewRecorder()
	err := j.Render(rr, req)
	if err != nil {
		t.Fatal(err)
	}
}
