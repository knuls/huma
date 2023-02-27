package mock

import (
	"context"
	"errors"

	"github.com/knuls/huma/pkg/creator"
)

type CreatorService struct{}

func NewCreatorService() *CreatorService {
	return &CreatorService{}
}

func (s *CreatorService) Find(ctx context.Context) ([]*creator.Creator, error) {
	return MockCreators, nil
}

func (s *CreatorService) FindById(ctx context.Context, id string) (*creator.Creator, error) {
	return MockCreators[0], nil
}

type ErrCreatorService struct{}

func NewErrCreatorService() *ErrCreatorService {
	return &ErrCreatorService{}
}

func (s *ErrCreatorService) Find(ctx context.Context) ([]*creator.Creator, error) {
	return nil, errors.New("some mock error")
}

func (s *ErrCreatorService) FindById(ctx context.Context, id string) (*creator.Creator, error) {
	return nil, errors.New("some mock error")
}
