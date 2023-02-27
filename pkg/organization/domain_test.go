package organization_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/knuls/huma/pkg/organization"
)

func TestOrganizationRender(t *testing.T) {
	user := organization.NewOrganization()
	req := &http.Request{}
	rr := httptest.NewRecorder()
	err := user.Render(rr, req)
	if err != nil {
		t.Fatal(err)
	}
}
