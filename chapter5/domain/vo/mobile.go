package vo

import (
	"errors"
	"github.com/tidwall/gjson"
	"regexp"
)

type Mobile struct {
	Value string `binding:"number,mobile"`
	Set   bool
}

func (o *Mobile) IsPresent(f func(v string)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Mobile) SetTo(v string) {
	o.Set = true
	o.Value = v
}

func (o *Mobile) Validate() error {
	if !regexp.MustCompile(`^(1[3-9][0-9]\d{8})$`).MatchString(o.Value) {
		return errors.New("手机号码格式不正确")
	}
	return nil
}

func (o *Mobile) UnmarshalJSON(data []byte) error {
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
