package user

import (
	"errors"
	"github.com/tidwall/gjson"
	"goODD/chapter6/domain/vo"
)

type Amount struct {
	vo.Money
	Value int64
	Set   bool
}

func (o *Amount) CashLimit() error {
	if o.Value < 100 {
		return errors.New("最低提现金额为1元")
	}
	return nil
}

func (o *Amount) Sufficient(v int64) bool {
	return o.Value >= v
}

func (o *Amount) Insufficient(v int64) bool {
	return !o.Sufficient(v)
}

func (o *Amount) ConvertFen(v float64) int64 {
	return o.Money.ConvertFen(v)
}

func (o *Amount) ConvertYuan() float64 {
	return o.Money.ConvertYuan(o.Value)
}

func (o *Amount) IsPresent(f func(v int64)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Amount) SetTo(v int64) {
	o.Set = true
	o.Value = v
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
