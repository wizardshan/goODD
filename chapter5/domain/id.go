package domain

import (
	"encoding/json"
	"github.com/tidwall/gjson"
)

type ID struct {
	Value int64 `binding:"min=1"`
	Set   bool
}

func (id *ID) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		result := gjson.GetBytes(data, "Value")
		if result.Exists() {
			id.Value = result.Int()
			id.Set = true
		}
		return nil
	}

	id.Value = gjson.ParseBytes(data).Int()
	id.Set = true
	return nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.Value)
}
