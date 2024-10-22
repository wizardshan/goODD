package vo

import "github.com/tidwall/gjson"

type ID struct {
	Value int64 `binding:"min=1"`
	Set   bool
}

func (o *ID) IsPresent(f func(v int64)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *ID) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

func (o *ID) UnmarshalJSON(data []byte) error {
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
