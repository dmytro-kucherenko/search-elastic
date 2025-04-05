package schema

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func Register(validate *validator.Validate) error {
	err := validate.RegisterValidation("query", ValidateQuery)
	if err != nil {
		return errors.New("validation registration error")
	}

	return nil
}

func New() (*validator.Validate, error) {
	validate := validator.New()
	err := Register(validate)

	return validate, err
}
