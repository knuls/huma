package creator_test

import (
	"context"
	"testing"
	"time"

	core "github.com/knuls/huma/pkg/core/dao"
	"github.com/knuls/huma/pkg/creator"
	"github.com/knuls/huma/pkg/creator/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestServiceFind(t *testing.T) {
	dao := mock.NewCreatorDao()
	errDao := mock.NewErrCreatorDao()
	tests := []struct {
		name string
		dao  core.Dao[creator.Creator]
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
			svc := creator.NewService(test.dao)
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
	dao := mock.NewCreatorDao()
	errDao := mock.NewErrCreatorDao()
	id := primitive.NewObjectIDFromTimestamp(time.Now()).Hex()

	tests := []struct {
		name string
		dao  core.Dao[creator.Creator]
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
			svc := creator.NewService(test.dao)
			_, err := svc.FindById(context.Background(), test.id)
			if err != nil {
				if !test.err {
					t.Error(err)
				}
			}
		})
	}
}
