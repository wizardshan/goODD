package user

import (
	"goODD/chapter7/domain/vo"
	"goODD/chapter7/pkg/validate"
)

type Avatar struct {
	vo.StringValue
}

func (o *Avatar) SetToDefault() {
	o.SetTo(o.Default())
}

func (o *Avatar) Default() string {
	return "avatar_default.png"
}

func (o *Avatar) ValidateOmit() error {
	if o.Set {
		return o.Validate()
	}
	return nil
}

func (o *Avatar) Validate() error {
	return validate.Var(o.Value, "image")
}
