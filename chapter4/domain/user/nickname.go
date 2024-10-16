package user

import "github.com/tidwall/gjson"

type Nickname struct {
	Value   string `binding:"min=2,max=10"`
	Keyword string `binding:"required"`
}

func (o *Nickname) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.Value = gjson.ParseBytes(data).String()
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Keyword")
	if results[0].Exists() {
		o.Value = results[0].String()
		return nil
	}
	if results[1].Exists() {
		o.Keyword = results[1].String()
	}
	return nil
}
