package validator

import (
	"testing"

	"github.com/go-playground/validator"
)

type test struct {
	Name string `json:"name" validate:"required"`
}

func TestValidator(t *testing.T) {
	v, err := New()
	if err != nil {
		t.Error(err)
	}
	o := &test{}
	o.Name = "some name"
	err = v.ValidateStruct(o)
	if err != nil {
		t.Error(err)
	}
}

func TestFailValidator(t *testing.T) {
	v, err := New()
	if err != nil {
		t.Error(err)
	}
	o := &test{}
	err = v.ValidateStruct(o)
	if err == nil {
		t.Error(err)
	}
}

func TestFailAddValidator(t *testing.T) {
	v, err := New()
	if err != nil {
		t.Error(err)
	}
	// key cannot be empty
	err = v.AddValidator("", func(fl validator.FieldLevel) bool {
		return false
	})
	if err == nil {
		t.Error(err)
	}
}
