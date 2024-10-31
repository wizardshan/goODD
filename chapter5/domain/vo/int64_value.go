package vo

import (
	"github.com/tidwall/gjson"
)

type Int64Value struct {
	Value int64
	Set   bool
}

func (o *Int64Value) IsPresent(f func(v int64)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Int64Value) Reset() {
	var v int64
	o.Value = v
	o.Set = false
}

func (o *Int64Value) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

func (o *Int64Value) Or(d int64) int64 {
	if o.Set {
		return o.Value
	}
	return d
}

func (o *Int64Value) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.SetTo(gjson.ParseBytes(data).Int())
		return nil
	}

	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		o.SetTo(result.Int())
	}
	return nil
}
