package user

import (
	"errors"
	"github.com/tidwall/gjson"
	"goODD/chapter10/domain/vo"
	"goODD/chapter10/pkg/validator"
)

type Amount struct {
	vo.Int64Range
	vo.Money
}

func (o Amount) ConvertFen(v float64) int64 {
	return o.Money.ConvertFen(v)
}

func (o Amount) ConvertYuan() float64 {
	return o.Money.ConvertYuan(o.Value)
}

func (o *Amount) CashLimit() error {
	if o.Value < 100 {
		return errors.New("最低提现金额为1元")
	}
	return nil
}

func (o Amount) Sufficient(v int64) bool {
	return o.Value >= v
}

func (o Amount) Insufficient(v int64) bool {
	return !o.Sufficient(v)
}

func (o Amount) validate(v int64) error {
	return validator.Var(v, "min=1,max=120")
}

func (o Amount) Validate() error {
	return o.validate(o.Value)
}

func (o Amount) ValidateOmit() error {
	if o.Set {
		return o.validate(o.Value)
	}
	return nil
}

func (o Amount) RangeValidate() error {
	return o.Int64Range.RangeValidate(o.validate)
}

func (o Amount) RangeValidateOmit() error {
	return o.Int64Range.RangeValidateOmit(o.validate)
}

func (o *Amount) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.SetTo(o.ConvertFen(gjson.ParseBytes(data).Float()))
		return nil
	}

	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		o.SetTo(o.ConvertFen(result.Float()))
	}
	return nil
}
