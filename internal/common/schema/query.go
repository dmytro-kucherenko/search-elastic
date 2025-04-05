package schema

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateTerms(terms string) bool {
	schema := regexp.MustCompile(`^[a-z ]+$`)

	return schema.MatchString(terms)
}

func ValidateQuery(field validator.FieldLevel) bool {
	return ValidateTerms(field.Field().String())
}
