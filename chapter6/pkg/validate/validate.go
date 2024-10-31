package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var validate = binding.Validator.Engine().(*validator.Validate)

func Var(field interface{}, tag string) error {
	return validate.Var(field, tag)
}
