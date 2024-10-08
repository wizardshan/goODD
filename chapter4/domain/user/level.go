package user

import (
	"errors"
	"github.com/tidwall/gjson"
)

type Level struct {
	Value int64 `binding:"oneof=0 10 20 30"`
	Set   bool
}

func (level *Level) UnmarshalJSON(data []byte) error {
	result := gjson.GetBytes(data, "Value")
	if !result.Exists() {
		return errors.New("`Value` key not found")
	}
	level.Value = result.Int()
	level.Set = true
	return nil
}
