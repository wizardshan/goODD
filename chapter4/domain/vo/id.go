package vo

import "github.com/tidwall/gjson"

type ID struct {
	Value int64 `binding:"min=1"`
}

func (o *ID) UnmarshalJSON(data []byte) error {
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
