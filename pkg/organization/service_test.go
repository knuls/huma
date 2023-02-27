package organization_test

import (
	"context"
	"testing"
	"time"

	core "github.com/knuls/huma/pkg/core/dao"
	"github.com/knuls/huma/pkg/organization"
	"github.com/knuls/huma/pkg/organization/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestServiceFind(t *testing.T) {
	dao := mock.NewOrganizationDao()
	errDao := mock.NewErrOrganizationDao()
	tests := []struct {
		name string
		dao  core.Dao[organization.Organization]
		err  bool
	}{
		{
			name: "findErr",
			dao:  errDao,
			err:  true,
		},
		{
			name: "findSuccess",
			dao:  dao,
			err:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			svc := organization.NewService(test.dao)
			_, err := svc.Find(context.Background())
			if err != nil {
				if !test.err {
					t.Error(err)
				}
			}
		})
	}
}

func TestServiceFindById(t *testing.T) {
	dao := mock.NewOrganizationDao()
	errDao := mock.NewErrOrganizationDao()
	id := primitive.NewObjectIDFromTimestamp(time.Now()).Hex()

	tests := []struct {
		name string
		dao  core.Dao[organization.Organization]
		err  bool
		id   string
	}{
		{
			name: "findByIdErr",
			dao:  errDao,
			err:  true,
			id:   id,
		},
		{
			name: "findByIdObjectIdErr",
			dao:  errDao,
			err:  true,
			id:   "some-bad-object-id",
		},
		{
			name: "findByIdSuccess",
			dao:  dao,
			err:  false,
			id:   id,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			svc := organization.NewService(test.dao)
			_, err := svc.FindById(context.Background(), test.id)
			if err != nil {
				if !test.err {
					t.Error(err)
				}
			}
		})
	}
}
