package creator_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/knuls/huma/pkg/creator"
	"github.com/knuls/huma/pkg/creator/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMux(t *testing.T) {
	t.Parallel()
	svc := mock.NewCreatorService()
	errSvc := mock.NewErrCreatorService()
	id := primitive.NewObjectIDFromTimestamp(time.Now())
	url := fmt.Sprintf("/%s", id.Hex())

	tests := []struct {
		name           string
		method         string
		path           string
		svc            creator.Service
		err            bool
		wantStatusCode int
	}{
		{
			name:           "findErrCreator",
			method:         http.MethodGet,
			path:           "/",
			svc:            errSvc,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "findCreator",
			method:         http.MethodGet,
			path:           "/",
			svc:            svc,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "findByIdErrCreator",
			method:         http.MethodGet,
			path:           url,
			svc:            errSvc,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "findbyIdCreator",
			method:         http.MethodGet,
			path:           url,
			svc:            svc,
			wantStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mux := creator.NewMux(test.svc)
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
