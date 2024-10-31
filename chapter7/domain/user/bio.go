package user

import (
	"goODD/chapter7/domain/vo"
	"goODD/chapter7/pkg/validate"
)

type Bio struct {
	vo.StringValue
}

func (o *Bio) ValidateOmit() error {
	if o.Set {
		return o.Validate()
	}
	return nil
}

func (o *Bio) Validate() error {
	return validate.Var(o.Value, "max=200")
}
