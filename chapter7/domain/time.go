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
	results := gjson.GetManyBytes(data, "Value", "Mask")
	if results[0].Exists() {
		t.Value = results[0].Time()
		t.Set = true
	}
	if results[1].Exists() {
		t.Mask = results[1].Bool()
	}
	return nil
}
