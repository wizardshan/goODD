package user

import (
	"github.com/tidwall/gjson"
)

type Bio struct {
	Value string
	Set   bool
}

func NewBio(v string) Bio {
	return Bio{Value: v, Set: true}
}

func (bio Bio) IsSet() bool {
	return bio.Set
}

func (bio Bio) IsPresent(f func(v string)) {
	if bio.Set {
		f(bio.Value)
	}
}

func (bio *Bio) UnmarshalJSON(data []byte) error {
	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		bio.Value = result.String()
		bio.Set = true
	}
	return nil
}
