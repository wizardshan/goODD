package user

import (
	"errors"
	"github.com/tidwall/gjson"
	"goODD/chapter6/pkg/validator"
)

type Age struct {
	Value int64
	Start int64
	End   int64
}

func (o Age) validate(v int64) error {
	return validator.Var(v, "min=1,max=120")
}

func (o Age) Validate() error {
	return o.validate(o.Value)
}

func (o Age) ValidateOmit() error {
	if o.Value != 0 {
		return o.validate(o.Value)
	}
	return nil
}

func (o Age) RangeValidate() error {
	if o.Start == 0 || o.End == 0 {
		return errors.New("开始或结束不能为空")
	}
	if err := o.validate(o.Start); err != nil {
		return err
	}
	if err := o.validate(o.End); err != nil {
		return err
	}
	if !o.StartBeforeEnd() {
		return errors.New("开始必须小于结束")
	}
	return nil
}

func (o Age) RangeValidateOmit() error {
	if o.Start != 0 {
		if err := o.validate(o.Start); err != nil {
			return err
		}
	}
	if o.End != 0 {
		if err := o.validate(o.End); err != nil {
			return err
		}
	}
	if o.Start != 0 && o.End != 0 && !o.StartBeforeEnd() {
		return errors.New("开始必须小于结束")
	}
	return nil
}

func (o Age) StartBeforeEnd() bool {
	return o.Start < o.End
}

func (o *Age) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.Value = gjson.ParseBytes(data).Int()
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Start", "End")
	if results[0].Exists() {
		o.Value = gjson.ParseBytes(data).Int()
		return nil
	}
	if results[1].Exists() {
		o.Start = results[1].Int()
	}
	if results[2].Exists() {
		o.End = results[2].Int()
	}
	return nil
}
