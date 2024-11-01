package vo

import "github.com/tidwall/gjson"

type StringValue struct {
	Value string
	Set   bool
}

func (o *StringValue) SetTo(v string) {
	o.Set = true
	o.Value = v
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
