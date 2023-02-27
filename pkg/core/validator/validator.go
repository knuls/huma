package validator

import (
	val "github.com/go-playground/validator"
	"github.com/knuls/huma/pkg/core/validator/validators"
)

var customValidators = map[string]val.Func{
	"oid": validators.ValidateObjectID,
}

type Validator struct {
	validate *val.Validate
}

func (v *Validator) ValidateStruct(o interface{}) error {
	return v.validate.Struct(o)
}

func (v *Validator) AddValidator(key string, fn val.Func) error {
	if err := v.validate.RegisterValidation(key, fn); err != nil {
		return err
	}
	return nil
}

func New() (*Validator, error) {
	validator := &Validator{validate: val.New()}
	for key, fn := range customValidators {
		validator.AddValidator(key, fn)
	}
	return validator, nil
}
