package creator_test

import (
	"context"
	"testing"
	"time"

	core "github.com/knuls/huma/pkg/core/dao"
	"github.com/knuls/huma/pkg/core/validator"
	"github.com/knuls/huma/pkg/creator"
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
			name: "findErrCreatorsNotFound",
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
			dao := creator.NewDao(t.Client, validator)
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
			name: "findOneErrCreatorsNotFound",
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
			dao := creator.NewDao(t.Client, validator)
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
	c := &creator.Creator{
		ID:        primitive.NewObjectID(),
		Email:     "some@email.com",
		FirstName: "some-first-name",
		LastName:  "some-last-name",
		Password:  "some-password",
		Verified:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tests := []struct {
		name    string
		mocks   []bson.D
		err     bool
		creator *creator.Creator
	}{
		{
			name: "createErrValidateStruct",
			mocks: []bson.D{
				{{Key: "ok", Value: 0}},
			},
			err:     true,
			creator: creator.NewCreator(),
		},
		{
			name: "createErrCreatorDuplicate",
			mocks: []bson.D{
				mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}),
				mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch),
			},
			err:     true,
			creator: c,
		},
		{
			name: "createSuccess",
			mocks: []bson.D{
				mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch),
				mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch),
				mtest.CreateSuccessResponse(),
			},
			err:     false,
			creator: c,
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
			dao := creator.NewDao(t.Client, validator)
			t.AddMockResponses(test.mocks...)
			_, err := dao.Create(context.Background(), test.creator)
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
			dao := creator.NewDao(t.Client, validator)
			t.AddMockResponses(test.mocks...)
			creator := &creator.Creator{
				Verified: false,
			}
			_, err := dao.Update(context.Background(), creator)
			if err == nil {
				t.Error(err)
			}
		})
	}
}
