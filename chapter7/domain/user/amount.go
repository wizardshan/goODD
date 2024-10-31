package user

import (
	"errors"
	"goODD/chapter7/domain/vo"
)

type Amount struct {
	vo.Int64Value
	vo.Money
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
