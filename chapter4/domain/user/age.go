package user

import "github.com/tidwall/gjson"

type Age struct {
	Value int64 `binding:"min=1,max=120"`
	Start int64 `binding:"min=1,max=120"`
	End   int64 `binding:"min=1,max=120"`
}

func (o *Age) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.Value = gjson.ParseBytes(data).Int()
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Start", "End")
	if results[0].Exists() {
		o.Value = gjson.ParseBytes(data).Int()
		return nil
	}
	if results[1].Exists() {
		o.Start = results[1].Int()
	}
	if results[2].Exists() {
		o.End = results[2].Int()
	}
	return nil
}
