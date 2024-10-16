package user

import (
	"goODD/chapter8/domain/vo"
	"goODD/chapter8/pkg/validator"
)

type Age struct {
	vo.Int64Range
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
