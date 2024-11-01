package validate

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Var(field interface{}, tag string) error {
	return validate.Var(field, tag)
}
