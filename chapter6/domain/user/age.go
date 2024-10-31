package user

import (
	"goODD/chapter5/domain/vo"
	"goODD/chapter5/pkg/validate"
)

type Age struct {
	vo.Int64Value
}

func (o *Age) ValidateOmit() error {
	if o.Set {
		return o.Validate()
	}
	return nil
}

func (o *Age) Validate() error {
	return validate.Var(o.Value, "min=1,max=120")
}
