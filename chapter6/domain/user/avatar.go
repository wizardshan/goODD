package user

import (
	"github.com/tidwall/gjson"
)

type Avatar struct {
	Value string `binding:"image"`
	Set   bool
}

func (o *Avatar) SetToDefault() {
	o.SetTo(o.Default())
}

func (o *Avatar) Default() string {
	return "avatar_default.png"
}

func (o *Avatar) IsPresent(f func(v string)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Avatar) SetTo(v string) {
	o.Set = true
	o.Value = v
}

func (o *Avatar) UnmarshalJSON(data []byte) error {
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
