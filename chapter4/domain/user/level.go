package user

import "github.com/tidwall/gjson"

type Level struct {
	Value int64 `binding:"oneof=0 10 20 30"`
	Set   bool
}

func (o *Level) IsPresent(f func(v int64)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Level) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

func (o *Level) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.Value = gjson.ParseBytes(data).Int()
		o.Set = true
		return nil
	}

	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		o.Value = result.Int()
		o.Set = true
	}
	return nil
}
