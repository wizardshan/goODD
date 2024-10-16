package vo

import (
	"errors"
	"github.com/tidwall/gjson"
)

type Int64Multi struct {
	Int64Value
	Values []int64
}

func (o Int64Multi) MultiEmpty() bool {
	return len(o.Values) == 0
}

func (o Int64Multi) MultiNotEmpty() bool {
	return !o.MultiEmpty()
}

func (o Int64Multi) MultiIsPresent(f func(v []int64)) {
	if o.MultiNotEmpty() {
		f(o.Values)
	}
}

func (o Int64Multi) MultiValidate(validate func(v int64) error) error {
	if o.MultiEmpty() {
		return errors.New("多值不能为空")
	}
	for _, v := range o.Values {
		if err := validate(v); err != nil {
			return err
		}
	}
	return nil
}

func (o Int64Multi) MultiValidateOmit(validate func(v int64) error) error {
	if o.MultiEmpty() {
		return nil
	}
	for _, v := range o.Values {
		if err := validate(v); err != nil {
			return err
		}
	}
	return nil
}

func (o *Int64Multi) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.SetTo(gjson.ParseBytes(data).Int())
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Values")
	if results[0].Exists() {
		o.SetTo(results[0].Int())
		return nil
	}
	if results[1].Exists() {
		values := results[1].Array()
		for _, s := range values {
			o.Values = append(o.Values, s.Int())
		}
	}
	return nil
}
