package user

import (
	"goODD/chapter7/domain/vo"
	"goODD/chapter7/pkg/validator"
)

type Age struct {
	vo.Int64Range
}

func NewAge(v int64) Age {
	var o Age
	o.SetTo(v)
	return o
}

func (o Age) validate(v int64) error {
	return validator.Var(v, "min=1,max=120")
}

func (o Age) Validate() error {
	return o.validate(o.Value)
}

func (o Age) ValidateOmit() error {
	if o.Set {
		return o.validate(o.Value)
	}
	return nil
}

func (o Age) RangeValidate() error {
	return o.Int64Range.RangeValidate(o.validate)
}

func (o Age) RangeValidateOmit() error {
	return o.Int64Range.RangeValidateOmit(o.validate)
}
