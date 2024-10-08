package user

import (
	"encoding/json"
	"github.com/tidwall/gjson"
)

type Nickname struct {
	Value string
	Set   bool
}

func (nickname *Nickname) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		result := gjson.GetBytes(data, "Value")
		if result.Exists() {
			nickname.Value = result.String()
			nickname.Set = true
		}
		return nil
	}

	nickname.Value = gjson.ParseBytes(data).String()
	nickname.Set = true
	return nil
}

func (nickname Nickname) MarshalJSON() ([]byte, error) {
	return json.Marshal(nickname.Value)
}
