package organization

import (
	"context"

	core "github.com/knuls/huma/pkg/core/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Find(ctx context.Context) ([]*Organization, error)
	FindById(ctx context.Context, id string) (*Organization, error)
}

type svc struct {
	dao core.Dao[Organization]
}

func NewService(dao core.Dao[Organization]) *svc {
	return &svc{
		dao: dao,
	}
}

func (s *svc) Find(ctx context.Context) ([]*Organization, error) {
	orgs, err := s.dao.Find(ctx, core.Where{})
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

func (s *svc) FindById(ctx context.Context, id string) (*Organization, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	org, err := s.dao.FindOne(ctx, core.Where{{Key: "id", Value: oid}})
	if err != nil {
		return nil, err
	}
	return org, nil
}
