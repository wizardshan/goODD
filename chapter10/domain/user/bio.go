package user

import (
	"goODD/chapter10/domain/vo"
	"goODD/chapter10/pkg/validator"
)

type Bio struct {
	vo.StringValue
}

func (o Bio) validate(v string) error {
	return validator.Var(v, "max=200")
}

func (o Bio) Validate() error {
	return o.validate(o.Value)
}

func (o Bio) ValidateOmit() error {
	if o.Set {
		return o.validate(o.Value)
	}
	return nil
}
