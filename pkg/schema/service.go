package schema

import (
	"context"

	core "github.com/knuls/huma/pkg/core/dao"
)

type Service interface {
	Find(ctx context.Context) ([]*Schema, error)
}

type svc struct {
	dao core.Dao[Schema]
}

func NewService(dao core.Dao[Schema]) *svc {
	return &svc{
		dao: dao,
	}
}

func (s *svc) Find(ctx context.Context) ([]*Schema, error) {
	service, err := s.dao.Find(ctx, core.Where{})
	if err != nil {
		return nil, err
	}
	return service, nil
}
