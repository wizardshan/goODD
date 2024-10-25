package user

import "github.com/tidwall/gjson"

type Age struct {
	Value int64 `binding:"min=1,max=120"`
	Set   bool
}

func (o *Age) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

func (o *Age) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.SetTo(gjson.ParseBytes(data).Int())
		return nil
	}

	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		o.SetTo(result.Int())
	}
	return nil
}
