package organization_test

import (
	"context"
	"testing"
	"time"

	core "github.com/knuls/huma/pkg/core/dao"
	"github.com/knuls/huma/pkg/core/validator"
	"github.com/knuls/huma/pkg/organization"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestDaoFind(t *testing.T) {
	tests := []struct {
		name  string
		mocks []bson.D
		err   bool
	}{
		{
			name: "findErrOrganizationsNotFound",
			mocks: []bson.D{
				{{Key: "ok", Value: 0}},
			},
			err: true,
		},
		{
			name: "findSuccess",
			mocks: []bson.D{
				mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}),
				mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{}),
				mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch),
			},
			err: false,
		},
	}

	validator, err := validator.New()
	if err != nil {
		t.Error(err)
	}

	for _, test := range tests {
		mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
		defer mt.Close()
		mt.Run(test.name, func(t *mtest.T) {
			dao := organization.NewDao(t.Client, validator)
			t.AddMockResponses(test.mocks...)
			_, err := dao.Find(context.Background(), core.Where{})
			if err != nil {
				if !test.err {
					t.Error(err)
				}
			}
		})
	}
}

func TestDaoFindOne(t *testing.T) {
	tests := []struct {
		name  string
		mocks []bson.D
		err   bool
	}{
		{
			name: "findOneErrOrganizationsNotFound",
			mocks: []bson.D{
				{{Key: "ok", Value: 0}},
			},
			err: true,
		},
		{
			name: "findOneSuccess",
			mocks: []bson.D{
				mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}),
			},
			err: false,
		},
	}

	validator, err := validator.New()
	if err != nil {
		t.Error(err)
	}

	for _, test := range tests {
		mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
		defer mt.Close()
		mt.Run(test.name, func(t *mtest.T) {
			dao := organization.NewDao(t.Client, validator)
			t.AddMockResponses(test.mocks...)
			_, err := dao.FindOne(context.Background(), core.Where{})
			if err != nil {
				if !test.err {
					t.Error(err)
				}
			}
		})
	}
}

func TestDaoCreate(t *testing.T) {
	entity := &organization.Organization{
		ID:        primitive.NewObjectID(),
		Name:      "someorgname",
		CreatorID: primitive.NewObjectIDFromTimestamp(time.Now()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tests := []struct {
		name   string
		mocks  []bson.D
		err    bool
		entity *organization.Organization
	}{
		{
			name: "createErrValidateStruct",
			mocks: []bson.D{
				{{Key: "ok", Value: 0}},
			},
			err:    true,
			entity: organization.NewOrganization(),
		},
		{
			name: "createErrOrganizationDuplicate",
			mocks: []bson.D{
				mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}),
				mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch),
			},
			err:    true,
			entity: entity,
		},
		{
			name: "createSuccess",
			mocks: []bson.D{
				mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch),
				mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch),
				mtest.CreateSuccessResponse(),
			},
			err:    false,
			entity: entity,
		},
	}

	validator, err := validator.New()
	if err != nil {
		t.Error(err)
	}

	for _, test := range tests {
		mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
		defer mt.Close()
		mt.Run(test.name, func(t *mtest.T) {
			dao := organization.NewDao(t.Client, validator)
			t.AddMockResponses(test.mocks...)
			_, err := dao.Create(context.Background(), test.entity)
			if err != nil {
				if !test.err {
					t.Error(err)
				}
			}
		})
	}
}

func TestDaoUpdate(t *testing.T) {
	tests := []struct {
		name  string
		mocks []bson.D
	}{
		{
			name:  "updateSuccess",
			mocks: []bson.D{},
		},
	}

	validator, err := validator.New()
	if err != nil {
		t.Error(err)
	}

	for _, test := range tests {
		mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
		defer mt.Close()
		mt.Run(test.name, func(t *mtest.T) {
			dao := organization.NewDao(t.Client, validator)
			t.AddMockResponses(test.mocks...)
			organization := &organization.Organization{}
			_, err := dao.Update(context.Background(), organization)
			if err == nil {
				t.Error(err)
			}
		})
	}
}
