package validators

import (
	"testing"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type test struct {
	Name interface{} `validate:"oid"`
}

func TestValidateObjectIDError(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("oid", ValidateObjectID)
	s := test{
		Name: "name",
	}
	err := validate.Struct(s)
	if err == nil {
		t.Error(err, "should fail to convert string to object id")
	}
}

func TestValidateObjectID(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("oid", ValidateObjectID)
	s := test{
		Name: primitive.NilObjectID,
	}
	err := validate.Struct(s)
	if err != nil {
		t.Error(err)
	}
}
