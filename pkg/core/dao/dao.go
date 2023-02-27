package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Where bson.D

type finder[T any] interface {
	Find(ctx context.Context, filter Where) ([]*T, error)
	FindOne(ctx context.Context, filter Where) (*T, error)
}

type creator[T any] interface {
	Create(ctx context.Context, t *T) (string, error)
}

type updater[T any] interface {
	Update(ctx context.Context, t *T) (*T, error)
}

type Dao[T any] interface {
	finder[T]
	creator[T]
	updater[T]
}
