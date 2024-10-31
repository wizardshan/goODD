package user

import (
	"goODD/chapter5/domain/vo"
	"goODD/chapter5/pkg/validate"
)

type Nickname struct {
	vo.StringValue
}

func (o *Nickname) ValidateOmit() error {
	if o.Set {
		return o.Validate()
	}
	return nil
}

func (o *Nickname) Validate() error {
	return validate.Var(o.Value, "min=2,max=10")
}
