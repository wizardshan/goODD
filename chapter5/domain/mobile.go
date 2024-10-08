package domain

import (
	"encoding/json"
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
	if data[0] == '{' {
		result := gjson.GetBytes(data, "Value")
		if result.Exists() {
			mobile.Value = result.String()
			mobile.Set = true
		}
		return nil
	}

	mobile.Value = gjson.ParseBytes(data).String()
	mobile.Set = true
	return nil
}

func (mobile Mobile) MarshalJSON() ([]byte, error) {
	return json.Marshal(mobile.Value)
}
