package vo

import "github.com/tidwall/gjson"

type StringValue struct {
	Value string
	Set   bool
}

func (o *StringValue) IsPresent(f func(v string)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *StringValue) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

func (o *StringValue) SetTo(v string) {
	o.Set = true
	o.Value = v
}

func (o *StringValue) Or(d string) string {
	if o.Set {
		return o.Value
	}
	return d
}

func (o *StringValue) UnmarshalJSON(data []byte) error {
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
