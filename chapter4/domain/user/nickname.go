package user

import "github.com/tidwall/gjson"

type Nickname struct {
	Value string `binding:"min=2,max=10"`
	Set   bool
}

func (o *Nickname) SetTo(v string) {
	o.Set = true
	o.Value = v
}

func (o *Nickname) UnmarshalJSON(data []byte) error {
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
