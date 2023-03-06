package organization

import (
	"context"
	"errors"

	core "github.com/knuls/huma/pkg/core/dao"
	"github.com/knuls/huma/pkg/core/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME = "organizations"

type dao struct {
	client    *mongo.Client
	validator *validator.Validator
}

func NewDao(client *mongo.Client, validator *validator.Validator) *dao {
	return &dao{
		client:    client,
		validator: validator,
	}
}

func (d *dao) Find(ctx context.Context, filter core.Where) ([]*Organization, error) {
	var orgs []*Organization
	cursor, err := d.client.Database("huma").Collection(COLLECTION_NAME).Find(ctx, filter)
	if err != nil {
		if errors.Is(err, ErrOrganizationsNotFound) {
			return orgs, nil
		}
		return nil, err
	}
	if err = cursor.All(ctx, &orgs); err != nil {
		return nil, err
	}
	return orgs, nil
}

func (d *dao) FindOne(ctx context.Context, filter core.Where) (*Organization, error) {
	result := d.client.Database("huma").Collection(COLLECTION_NAME).FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		if errors.Is(err, ErrOrganizationsNotFound) {
			return nil, ErrOrganizationNotFound
		}
		return nil, err
	}
	var org *Organization
	if err = result.Decode(&org); err != nil {
		return nil, err
	}
	return org, nil
}

func (d *dao) Create(ctx context.Context, org *Organization) (string, error) {
	if err := d.validator.ValidateStruct(org); err != nil {
		return "", err
	}
	exists, err := d.Find(ctx, core.Where{})
	if err != nil {
		return "", err
	}
	if len(exists) > 0 {
		return "", ErrOrganizationDuplicate
	}
	result, err := d.client.Database("huma").Collection(COLLECTION_NAME).InsertOne(ctx, org)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (d *dao) Update(ctx context.Context, org *Organization) (*Organization, error) {
	return nil, ErrNoImpl
}
