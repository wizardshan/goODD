package user

import (
	"chapter10/domain/vo"
	"chapter10/rpc/domain/user"
	"errors"
)

type Amount struct {
	vo.Money
	user.Amount
	Set bool
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

func (o *Amount) SetToPb(v *user.Amount) {
	if v != nil {
		o.SetTo(v.Value)
	}
}
