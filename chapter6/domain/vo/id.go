package vo

import (
	"errors"
	"github.com/tidwall/gjson"
)

type ID struct {
	Value int64
}

func (o ID) validate(v int64) error {
	if v < 1 {
		return errors.New("ID必须大于等于1")
	}
	return nil
}

func (o ID) Validate() error {
	return o.validate(o.Value)
}

func (o ID) ValidateOmit() error {
	if o.Value != 0 {
		return o.validate(o.Value)
	}
	return nil
}

func (o *ID) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.Value = gjson.ParseBytes(data).Int()
		return nil
	}

	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		o.Value = result.Int()
	}
	return nil
}
