package user

import (
	"goODD/chapter7/domain/vo"
	"goODD/chapter7/pkg/validate"
)

const (
	StatusNormal = 0
	StatusCancel = 10
	StatusForbid = 20
)

type Statuses []Status

func (o Statuses) IsPresent(f func(v []int64)) {
	if len(o) != 0 {
		f(o.Values())
	}
}

func (o Statuses) Values() []int64 {
	values := make([]int64, len(o))
	for i, status := range o {
		values[i] = status.Value
	}
	return values
}

type Status struct {
	vo.Int64Value
}

func (o *Status) IsCancel() bool {
	return o.Value == StatusCancel
}

func (o *Status) IsForbid() bool {
	return o.Value == StatusForbid
}

func (o *Status) Desc() string {
	mapping := map[int64]string{
		StatusNormal: "正常",
		StatusCancel: "注销",
		StatusForbid: "禁止",
	}
	desc, ok := mapping[o.Value]
	if ok {
		return desc
	}
	return "未知"
}

func (o *Status) ValidateOmit() error {
	if o.Set {
		return o.Validate()
	}
	return nil
}

func (o *Status) Validate() error {
	return validate.Var(o.Value, "oneof=0 10 20")
}
