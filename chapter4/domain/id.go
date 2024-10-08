package domain

import (
	"github.com/tidwall/gjson"
)

type ID struct {
	Value int64 `binding:"min=1"`
	Set   bool
}

func (id *ID) UnmarshalJSON(data []byte) error {
	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		id.Value = result.Int()
		id.Set = true
	}
	return nil
}
