package user

import (
	"goODD/chapter7/domain/vo"
	"goODD/chapter7/pkg/validator"
)

type Level struct {
	vo.Int64Multi
}

func NewLevel(v int64) Level {
	var o Level
	o.SetTo(v)
	return o
}
func (o Level) validate(v int64) error {
	return validator.Var(v, "oneof=0 10 20 30")
}

func (o Level) Validate() error {
	return o.validate(o.Value)
}

func (o Level) ValidateOmit() error {
	if o.Set {
		return o.validate(o.Value)
	}
	return nil
}

func (o Level) MultiValidate() error {
	return o.Int64Multi.MultiValidate(o.validate)
}

func (o Level) MultiValidateOmit() error {
	return o.Int64Multi.MultiValidateOmit(o.validate)
}
