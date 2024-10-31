package user

import (
	"goODD/chapter5/domain/vo"
	"goODD/chapter5/pkg/validate"
)

type Levels []Level

func (o Levels) IsPresent(f func(v []int64)) {
	if len(o) != 0 {
		f(o.Values())
	}
}

func (o Levels) Values() []int64 {
	values := make([]int64, len(o))
	for i, level := range o {
		values[i] = level.Value
	}
	return values
}

type Level struct {
	vo.Int64Value
}

func (o *Level) ValidateOmit() error {
	if o.Set {
		return o.Validate()
	}
	return nil
}

func (o *Level) Validate() error {
	return validate.Var(o.Value, "oneof=0 10 20 30")
}
