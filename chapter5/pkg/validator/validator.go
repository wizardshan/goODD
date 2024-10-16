package validator

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func Var(v any, tag string) error {
	return validate.Var(v, tag)
}
