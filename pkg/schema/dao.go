package schema

import (
	"context"
	"errors"

	core "github.com/knuls/huma/pkg/core/dao"
	"github.com/knuls/huma/pkg/core/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME = "schemas"

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

func (d *dao) Find(ctx context.Context, filter core.Where) ([]*Schema, error) {
	name, ok := ctx.Value(OrganizationNameCtxKey{}).(string)
	if !ok {
		return nil, errors.New("cast error")
	}
	var orgs []*Schema
	col := d.client.Database(name).Collection(COLLECTION_NAME)
	cursor, err := col.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, ErrSchemasNotFound) {
			return orgs, nil
		}
		return nil, err
	}
	if err = cursor.All(ctx, &orgs); err != nil {
		return nil, err
	}
	return orgs, nil
}

func (d *dao) FindOne(ctx context.Context, filter core.Where) (*Schema, error) {
	return nil, ErrNoImpl
}

func (d *dao) Create(ctx context.Context, org *Schema) (string, error) {
	return "", ErrNoImpl

}

func (d *dao) Update(ctx context.Context, org *Schema) (*Schema, error) {
	return nil, ErrNoImpl
}
