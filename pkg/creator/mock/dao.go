package mock

import (
	"context"
	"errors"
	"time"

	"github.com/knuls/huma/pkg/core/dao"
	"github.com/knuls/huma/pkg/creator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MockCreators = []*creator.Creator{
	{
		ID:        primitive.NewObjectIDFromTimestamp(time.Now()),
		Email:     "first@knuls.io",
		FirstName: "first",
		LastName:  "knuls",
		Password:  "super-secret",
		Verified:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        primitive.NewObjectIDFromTimestamp(time.Now().Add(5 * time.Minute)),
		Email:     "second@knuls.io",
		FirstName: "second",
		LastName:  "knuls",
		Password:  "super-secret",
		Verified:  true,
		CreatedAt: time.Now().Add(5 * time.Minute),
		UpdatedAt: time.Now().Add(5 * time.Minute),
	},
	{
		ID:        primitive.NewObjectIDFromTimestamp(time.Now().Add(10 * time.Minute)),
		Email:     "third@knuls.io",
		FirstName: "third",
		LastName:  "knusecols",
		Password:  "super-secret",
		Verified:  false,
		CreatedAt: time.Now().Add(10 * time.Minute),
		UpdatedAt: time.Now().Add(10 * time.Minute),
	},
}

type CreatorDao struct {
}

func NewCreatorDao() *CreatorDao {
	return &CreatorDao{}
}

func (m *CreatorDao) Find(ctx context.Context, filter dao.Where) ([]*creator.Creator, error) {
	return MockCreators, nil
}
func (m *CreatorDao) FindOne(ctx context.Context, filter dao.Where) (*creator.Creator, error) {
	return MockCreators[0], nil
}
func (m *CreatorDao) Create(ctx context.Context, creator *creator.Creator) (string, error) {
	return "", nil
}
func (m *CreatorDao) Update(ctx context.Context, creator *creator.Creator) (*creator.Creator, error) {
	return nil, nil
}

type ErrCreatorDao struct {
}

func NewErrCreatorDao() *ErrCreatorDao {
	return &ErrCreatorDao{}
}

func (m *ErrCreatorDao) Find(ctx context.Context, filter dao.Where) ([]*creator.Creator, error) {
	return nil, errors.New("some mock error")
}
func (m *ErrCreatorDao) FindOne(ctx context.Context, filter dao.Where) (*creator.Creator, error) {
	return nil, errors.New("some mock error")
}
func (m *ErrCreatorDao) Create(ctx context.Context, creator *creator.Creator) (string, error) {
	return "", nil
}
func (m *ErrCreatorDao) Update(ctx context.Context, creator *creator.Creator) (*creator.Creator, error) {
	return nil, nil
}
