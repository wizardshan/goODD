package vo

import (
	"github.com/tidwall/gjson"
)

type Int64Value struct {
	Value int64
	Set   bool
}

func (o *Int64Value) SetTo(v int64) {
	o.Set = true
	o.Value = v
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
