package validators

import (
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateObjectID(fl validator.FieldLevel) bool {
	id, ok := fl.Field().Interface().(primitive.ObjectID)
	if !ok {
		return false
	}
	return primitive.IsValidObjectID(id.Hex())
}
