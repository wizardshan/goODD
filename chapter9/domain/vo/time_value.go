package vo

import (
	"github.com/tidwall/gjson"
	"time"
)

type TimeValue struct {
	Value time.Time
	Set   bool
}

func (o TimeValue) IsPresent(f func(v time.Time)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *TimeValue) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

func (o *TimeValue) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

func (o TimeValue) Or(d time.Time) time.Time {
	if o.Set {
		return o.Value
	}
	return d
}

func (o *TimeValue) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.SetTo(gjson.ParseBytes(data).Time())
		return nil
	}

	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		o.SetTo(result.Time())
	}
	return nil
}
