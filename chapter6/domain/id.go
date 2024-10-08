package domain

import (
	"github.com/tidwall/gjson"
)

type ID struct {
	Value int64 `binding:"min=1"`
	Set   bool
	Mask  bool
}

func (id *ID) UnmarshalJSON(data []byte) error {
	results := gjson.GetManyBytes(data, "Value", "Mask")
	if results[0].Exists() {
		id.Value = results[0].Int()
		id.Set = true
	}
	if results[1].Exists() {
		id.Mask = results[1].Bool()
	}
	return nil
}
