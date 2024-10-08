package user

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
)

type Level struct {
	Value int64 `binding:"oneof=0 10 20 30"`
	Set   bool
}

func (level *Level) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		result := gjson.GetBytes(data, "Value")
		if !result.Exists() {
			return errors.New("`Value` key not found")
		}
		level.Value = result.Int()
		level.Set = true
		return nil
	}

	level.Value = gjson.ParseBytes(data).Int()
	level.Set = true
	return nil
}

func (level Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(level.Value)
}
