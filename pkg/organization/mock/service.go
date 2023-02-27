package mock

import (
	"context"
	"errors"

	"github.com/knuls/huma/pkg/organization"
)

type OrganizationService struct{}

func NewOrganizationService() *OrganizationService {
	return &OrganizationService{}
}

func (s *OrganizationService) Find(ctx context.Context) ([]*organization.Organization, error) {
	return MockOrganizations, nil
}

func (s *OrganizationService) FindById(ctx context.Context, id string) (*organization.Organization, error) {
	return MockOrganizations[0], nil
}

type ErrOrganizationService struct{}

func NewErrOrganizationService() *ErrOrganizationService {
	return &ErrOrganizationService{}
}

func (s *ErrOrganizationService) Find(ctx context.Context) ([]*organization.Organization, error) {
	return nil, errors.New("some mock error")
}

func (s *ErrOrganizationService) FindById(ctx context.Context, id string) (*organization.Organization, error) {
	return nil, errors.New("some mock error")
}
