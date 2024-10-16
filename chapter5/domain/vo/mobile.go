package vo

import (
	"errors"
	"github.com/tidwall/gjson"
	"regexp"
)

type Mobile struct {
	Value string
}

func (o Mobile) validate(v string) error {
	if !regexp.MustCompile(`^(1[3-9][0-9]\d{8})$`).MatchString(v) {
		return errors.New("手机号码格式不正确")
	}
	return nil
}

func (o Mobile) Validate() error {
	return o.validate(o.Value)
}

func (o Mobile) ValidateOmit() error {
	if o.Value != "" {
		return o.validate(o.Value)
	}
	return nil
}

func (o *Mobile) UnmarshalJSON(data []byte) error {
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
