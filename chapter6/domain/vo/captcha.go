package vo

import (
	"github.com/tidwall/gjson"
)

type Captcha struct {
	Value string `binding:"number,len=4"`
	Set   bool
}

func (o *Captcha) IsPresent(f func(v string)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Captcha) SetTo(v string) {
	o.Set = true
	o.Value = v
}

func (o *Captcha) UnmarshalJSON(data []byte) error {
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
