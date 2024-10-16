package vo

import (
	"errors"
)

type ID struct {
	Int64Value
}

func NewID(v int64) ID {
	var o ID
	o.SetTo(v)
	return o
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
	if o.Set {
		return o.validate(o.Value)
	}
	return nil
}
