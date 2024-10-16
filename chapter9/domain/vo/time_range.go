package vo

import (
	"errors"
	"github.com/tidwall/gjson"
	"time"
)

type TimeRange struct {
	TimeValue
	Start TimeValue
	End   TimeValue
}

func (o TimeRange) StartIsPresent(f func(v time.Time)) {
	if o.Start.Set {
		f(o.Start.Value)
	}
}

func (o TimeRange) EndIsPresent(f func(v time.Time)) {
	if o.End.Set {
		f(o.End.Value)
	}
}

func (o TimeRange) StartBeforeEnd() bool {
	return o.Start.Value.Before(o.End.Value)
}

func (o TimeRange) RangeValidate(validate func(v time.Time) error) error {
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

func (o TimeRange) RangeValidateOmit(validate func(v time.Time) error) error {
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

func (o *TimeRange) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.SetTo(gjson.ParseBytes(data).Time())
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Start", "End")
	if results[0].Exists() {
		o.SetTo(results[0].Time())
		return nil
	}
	if results[1].Exists() {
		o.Start.SetTo(results[1].Time())
	}
	if results[2].Exists() {
		o.End.SetTo(results[2].Time())
	}
	return nil
}
