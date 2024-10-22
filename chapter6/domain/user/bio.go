package user

import (
	"github.com/tidwall/gjson"
)

type Bio struct {
	Value string `binding:"max=200"`
	Set   bool
}

func (o *Bio) IsPresent(f func(v string)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Bio) SetTo(v string) {
	o.Set = true
	o.Value = v
}

func (o *Bio) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.SetTo(gjson.ParseBytes(data).String())
		return nil
	}

	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		o.SetTo(result.String())
	}
	return nil
}
