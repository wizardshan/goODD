package user

import (
	"goODD/chapter5/domain/vo"
	"goODD/chapter5/pkg/validate"
)

type Level struct {
	vo.Int64Value
}

func (o *Level) ValidateOmit() error {
	if o.Set {
		return o.Validate()
	}
	return nil
}

func (o *Level) Validate() error {
	return validate.Var(o.Value, "oneof=0 10 20 30")
}
