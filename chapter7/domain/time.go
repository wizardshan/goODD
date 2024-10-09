package domain

import (
	"github.com/tidwall/gjson"
	"time"
)

type Time struct {
	Value time.Time
	Set   bool
	Mask  bool
}

func NewTime(v time.Time) Time {
	return Time{Value: v, Set: true}
}

func (t Time) IsMask() bool {
	return t.Mask
}

func (t Time) IsSet() bool {
	return t.Set
}

func (t Time) IsPresent(f func(v time.Time)) {
	if t.Set {
		f(t.Value)
	}
}

func (t *Time) UnmarshalJSON(data []byte) error {
	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		t.Value = result.Time()
		t.Set = true
	}
	return nil
}
