package vo

import "github.com/tidwall/gjson"

type Captcha struct {
	Value string `binding:"number,len=4"`
}

func (o *Captcha) UnmarshalJSON(data []byte) error {
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
