package domain

import (
	"github.com/tidwall/gjson"
)

type ID struct {
	Value int64 `binding:"min=1"`
	Set   bool
}

func NewID(v int64) ID {
	return ID{Value: v, Set: true}
}

func (id ID) IsSet() bool {
	return id.Set
}

func (id ID) IsPresent(f func(v int64)) {
	if id.Set {
		f(id.Value)
	}
}

func (id *ID) UnmarshalJSON(data []byte) error {
	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		id.Value = result.Int()
		id.Set = true
	}
	return nil
}
