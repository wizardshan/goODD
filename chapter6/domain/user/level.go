package user

import (
	"errors"
	"github.com/tidwall/gjson"
	"goODD/chapter6/pkg/validator"
)

type Level struct {
	Value  int64
	Values []int64
}

func (o Level) validate(v int64) error {
	return validator.Var(v, "oneof=0 10 20 30")
}

func (o Level) Validate() error {
	return o.validate(o.Value)
}

func (o Level) ValidateOmit() error {
	if o.Value != 0 {
		return o.validate(o.Value)
	}
	return nil
}

func (o Level) MultiValidate() error {
	if len(o.Values) == 0 {
		return errors.New("多值不能为空")
	}
	for _, v := range o.Values {
		if err := o.validate(v); err != nil {
			return err
		}
	}
	return nil
}

func (o Level) MultiValidateOmit() error {
	if len(o.Values) == 0 {
		return nil
	}
	for _, v := range o.Values {
		if err := o.validate(v); err != nil {
			return err
		}
	}
	return nil
}

func (o *Level) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.Value = gjson.ParseBytes(data).Int()
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Values")
	if results[0].Exists() {
		o.Value = results[0].Int()
		return nil
	}
	if results[1].Exists() {
		values := results[1].Array()
		for _, v := range values {
			o.Values = append(o.Values, v.Int())
		}
	}
	return nil
}
