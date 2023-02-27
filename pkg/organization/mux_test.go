package organization_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/knuls/huma/pkg/organization"
	"github.com/knuls/huma/pkg/organization/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMux(t *testing.T) {
	t.Parallel()
	svc := mock.NewOrganizationService()
	errSvc := mock.NewErrOrganizationService()
	id := primitive.NewObjectIDFromTimestamp(time.Now())
	url := fmt.Sprintf("/%s", id.Hex())

	tests := []struct {
		name           string
		method         string
		path           string
		svc            organization.Service
		err            bool
		wantStatusCode int
	}{
		{
			name:           "findErrOrganization",
			method:         http.MethodGet,
			path:           "/",
			svc:            errSvc,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "findOrganization",
			method:         http.MethodGet,
			path:           "/",
			svc:            svc,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "findByIdErrOrganization",
			method:         http.MethodGet,
			path:           url,
			svc:            errSvc,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "findbyIdOrganization",
			method:         http.MethodGet,
			path:           url,
			svc:            svc,
			wantStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mux := organization.NewMux(test.svc)
			req := httptest.NewRequest(test.method, test.path, nil)
			rr := httptest.NewRecorder()
			mux.Routes().ServeHTTP(rr, req)
			res := rr.Result()
			if res.StatusCode != test.wantStatusCode {
				t.Fatalf("want reponse status code %d, got %d", test.wantStatusCode, res.StatusCode)
			}
		})
	}
}
