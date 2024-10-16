package user

import "github.com/tidwall/gjson"

type Level struct {
	Value int64 `binding:"oneof=0 10 20 30"`
}

func (o *Level) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.Value = gjson.ParseBytes(data).Int()
		return nil
	}

	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		o.Value = result.Int()
	}
	return nil
}
