package creator

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNoImpl           = errors.New("error no impl")
	ErrCreatorNotFound  = errors.New("error creator not found")
	ErrCreatorDuplicate = errors.New("error duplicate creator")
	ErrCreatorsNotFound = mongo.ErrNoDocuments
)
