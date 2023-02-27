package mock

import (
	"context"
	"errors"
	"time"

	"github.com/knuls/huma/pkg/core/dao"
	"github.com/knuls/huma/pkg/organization"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MockOrganizations = []*organization.Organization{
	{
		ID: primitive.NewObjectIDFromTimestamp(time.Now()),
	},
	{
		ID: primitive.NewObjectIDFromTimestamp(time.Now().Add(5 * time.Minute)),
	},
	{
		ID: primitive.NewObjectIDFromTimestamp(time.Now().Add(10 * time.Minute)),
	},
}

type OrganizationDao struct {
}

func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{}
}

func (m *OrganizationDao) Find(ctx context.Context, filter dao.Where) ([]*organization.Organization, error) {
	return MockOrganizations, nil
}
func (m *OrganizationDao) FindOne(ctx context.Context, filter dao.Where) (*organization.Organization, error) {
	return MockOrganizations[0], nil
}
func (m *OrganizationDao) Create(ctx context.Context, org *organization.Organization) (string, error) {
	return "", nil
}
func (m *OrganizationDao) Update(ctx context.Context, org *organization.Organization) (*organization.Organization, error) {
	return nil, nil
}

type ErrOrganizationDao struct {
}

func NewErrOrganizationDao() *ErrOrganizationDao {
	return &ErrOrganizationDao{}
}

func (m *ErrOrganizationDao) Find(ctx context.Context, filter dao.Where) ([]*organization.Organization, error) {
	return nil, errors.New("some mock error")
}
func (m *ErrOrganizationDao) FindOne(ctx context.Context, filter dao.Where) (*organization.Organization, error) {
	return nil, errors.New("some mock error")
}
func (m *ErrOrganizationDao) Create(ctx context.Context, org *organization.Organization) (string, error) {
	return "", nil
}
func (m *ErrOrganizationDao) Update(ctx context.Context, org *organization.Organization) (*organization.Organization, error) {
	return nil, nil
}
