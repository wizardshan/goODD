package vo

import (
	"errors"
	"github.com/tidwall/gjson"
)

type Int64Range struct {
	Int64Value
	Start Int64Value
	End   Int64Value
}

func (o *Int64Range) RangeSetTo(start int64, end int64) {
	o.Start.SetTo(start)
	o.End.SetTo(end)
}

func (o *Int64Range) StartSetTo(v int64) {
	o.Start.SetTo(v)
}

func (o *Int64Range) EndSetTo(v int64) {
	o.End.SetTo(v)
}

func (o Int64Range) StartIsPresent(f func(v int64)) {
	if o.Start.Set {
		f(o.Start.Value)
	}
}

func (o Int64Range) EndIsPresent(f func(v int64)) {
	if o.End.Set {
		f(o.End.Value)
	}
}

func (o Int64Range) StartBeforeEnd() bool {
	return o.Start.Value < o.End.Value
}

func (o Int64Range) RangeDiff() int64 {
	return o.End.Value - o.Start.Value
}

func (o Int64Range) RangeValidate(validate func(v int64) error) error {
	if !o.Start.Set || !o.End.Set {
		return errors.New("开始或结束不能为空")
	}
	if err := validate(o.Start.Value); err != nil {
		return err
	}
	if err := validate(o.End.Value); err != nil {
		return err
	}
	if !o.StartBeforeEnd() {
		return errors.New("开始必须小于结束")
	}
	return nil
}

func (o Int64Range) RangeValidateOmit(validate func(v int64) error) error {
	if o.Start.Set {
		if err := validate(o.Start.Value); err != nil {
			return err
		}
	}
	if o.End.Set {
		if err := validate(o.End.Value); err != nil {
			return err
		}
	}
	if o.Start.Set && o.End.Set && !o.StartBeforeEnd() {
		return errors.New("开始必须小于结束")
	}
	return nil
}

func (o *Int64Range) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.SetTo(gjson.ParseBytes(data).Int())
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Start", "End")
	if results[0].Exists() {
		o.SetTo(results[0].Int())
		return nil
	}
	if results[1].Exists() {
		o.Start.SetTo(results[1].Int())
	}
	if results[2].Exists() {
		o.End.SetTo(results[2].Int())
	}
	return nil
}
