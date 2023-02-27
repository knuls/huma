package creator_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/knuls/huma/pkg/creator"
)

func TestCreatorRender(t *testing.T) {
	user := creator.NewCreator()
	req := &http.Request{}
	rr := httptest.NewRecorder()
	err := user.Render(rr, req)
	if err != nil {
		t.Fatal(err)
	}
}
