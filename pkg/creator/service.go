package creator

import (
	"context"

	core "github.com/knuls/huma/pkg/core/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Find(ctx context.Context) ([]*Creator, error)
	FindById(ctx context.Context, id string) (*Creator, error)
}

type svc struct {
	dao core.Dao[Creator]
}

func NewService(dao core.Dao[Creator]) *svc {
	return &svc{
		dao: dao,
	}
}

func (s *svc) Find(ctx context.Context) ([]*Creator, error) {
	creators, err := s.dao.Find(ctx, core.Where{})
	if err != nil {
		return nil, err
	}
	return creators, nil
}

func (s *svc) FindById(ctx context.Context, id string) (*Creator, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	creator, err := s.dao.FindOne(ctx, core.Where{{Key: "id", Value: oid}})
	if err != nil {
		return nil, err
	}
	return creator, nil
}
