package user

import "github.com/tidwall/gjson"

type Age struct {
	Value int64 `binding:"min=1,max=120"`
}

func (o *Age) UnmarshalJSON(data []byte) error {
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
