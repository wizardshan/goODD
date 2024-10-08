package domain

import (
	"errors"
	"github.com/tidwall/gjson"
	"regexp"
)

type Mobile struct {
	Value string `binding:"number,mobile"`
	Set   bool
}

func (mobile Mobile) IsValid() bool {
	return regexp.MustCompile(`^(1[3-9][0-9]\d{8})$`).MatchString(mobile.Value)
}

func (mobile *Mobile) UnmarshalJSON(data []byte) error {
	result := gjson.GetBytes(data, "Value")
	if !result.Exists() {
		return errors.New("`Value` key not found")
	}
	mobile.Value = result.String()
	mobile.Set = true
	return nil
}
