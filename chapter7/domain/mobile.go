package domain

import (
	"github.com/tidwall/gjson"
	"regexp"
)

type Mobile struct {
	Value string `binding:"number,mobile"`
	Set   bool
}

func NewMobile(v string) Mobile {
	return Mobile{Value: v, Set: true}
}

func (mobile Mobile) IsSet() bool {
	return mobile.Set
}

func (mobile Mobile) IsPresent(f func(v string)) {
	if mobile.Set {
		f(mobile.Value)
	}
}

func (mobile Mobile) IsValid() bool {
	return regexp.MustCompile(`^(1[3-9][0-9]\d{8})$`).MatchString(mobile.Value)
}

func (mobile *Mobile) UnmarshalJSON(data []byte) error {
	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		mobile.Value = result.String()
		mobile.Set = true
	}
	return nil
}
