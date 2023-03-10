package creator

import (
	"context"
	"errors"
	"time"

	core "github.com/knuls/huma/pkg/core/dao"
	"github.com/knuls/huma/pkg/core/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type dao struct {
	creators  *mongo.Collection
	validator *validator.Validator
}

func NewDao(db *mongo.Client, validator *validator.Validator) *dao {
	return &dao{
		creators:  db.Database("huma").Collection("creators"),
		validator: validator,
	}
}

func (d *dao) Find(ctx context.Context, filter core.Where) ([]*Creator, error) {
	var creators []*Creator
	cursor, err := d.creators.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, ErrCreatorsNotFound) {
			return creators, nil
		}
		return nil, err
	}
	if err = cursor.All(ctx, &creators); err != nil {
		return nil, err
	}
	return creators, nil
}

func (d *dao) FindOne(ctx context.Context, filter core.Where) (*Creator, error) {
	result := d.creators.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		if errors.Is(err, ErrCreatorsNotFound) {
			return nil, ErrCreatorNotFound
		}
		return nil, err
	}
	var creator *Creator
	if err = result.Decode(&creator); err != nil {
		return nil, err
	}
	return creator, nil
}

func (d *dao) Create(ctx context.Context, creator *Creator) (string, error) {
	if err := d.validator.ValidateStruct(creator); err != nil {
		return "", err
	}
	exists, err := d.Find(ctx, core.Where{})
	if err != nil {
		return "", err
	}
	if len(exists) > 0 {
		return "", ErrCreatorDuplicate
	}
	now := time.Now()
	creator.Verified = false
	creator.CreatedAt = now
	creator.UpdatedAt = now
	result, err := d.creators.InsertOne(ctx, creator)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (d *dao) Update(ctx context.Context, creator *Creator) (*Creator, error) {
	return nil, ErrNoImpl
}
