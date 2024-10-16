package user

import "github.com/tidwall/gjson"

type Nickname struct {
	Value string `binding:"min=2,max=10"`
}

func (o *Nickname) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.Value = gjson.ParseBytes(data).String()
		return nil
	}

	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		o.Value = result.String()
	}
	return nil
}
