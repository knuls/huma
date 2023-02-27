package organization

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNoImpl                = errors.New("error no impl")
	ErrOrganizationNotFound  = errors.New("error organization not found")
	ErrOrganizationDuplicate = errors.New("error duplicate organization")
	ErrOrganizationsNotFound = mongo.ErrNoDocuments
)
