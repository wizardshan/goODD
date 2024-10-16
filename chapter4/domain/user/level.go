package user

import "github.com/tidwall/gjson"

type Level struct {
	Value  int64   `binding:"oneof=0 10 20 30"`
	Values []int64 `binding:"dive,oneof=0 10 20 30"`
}

func (o *Level) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.Value = gjson.ParseBytes(data).Int()
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Values")
	if results[0].Exists() {
		o.Value = results[0].Int()
		return nil
	}
	if results[1].Exists() {
		values := results[1].Array()
		for _, v := range values {
			o.Values = append(o.Values, v.Int())
		}
	}
	return nil
}
