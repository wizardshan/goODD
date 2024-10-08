package user

import (
	"github.com/tidwall/gjson"
)

type Level struct {
	Value int64 `binding:"oneof=0 10 20 30"`
	Set   bool
	Mask  bool
}

func (level *Level) UnmarshalJSON(data []byte) error {
	results := gjson.GetManyBytes(data, "Value", "Mask")
	if results[0].Exists() {
		level.Value = results[0].Int()
		level.Set = true
	}
	if results[1].Exists() {
		level.Mask = results[1].Bool()
	}
	return nil
}
