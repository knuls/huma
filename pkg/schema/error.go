package schema

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNoImpl          = errors.New("error no impl")
	ErrSchemaNotFound  = errors.New("error schema not found")
	ErrSchemaDuplicate = errors.New("error duplicate organization")
	ErrSchemasNotFound = mongo.ErrNoDocuments
)
