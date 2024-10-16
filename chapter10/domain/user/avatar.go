package user

import (
	"goODD/chapter10/domain/vo"
	"goODD/chapter10/pkg/validator"
)

type Avatar struct {
	vo.StringValue
}

func (o *Avatar) SetToDefault() {
	o.SetTo(o.Default())
}

func (o Avatar) Default() string {
	return "avatar_default.png"
}

func (o Avatar) validate(v string) error {
	return validator.Var(v, "image")
}

func (o Avatar) Validate() error {
	return o.validate(o.Value)
}

func (o Avatar) ValidateOmit() error {
	if o.Set {
		return o.validate(o.Value)
	}
	return nil
}
