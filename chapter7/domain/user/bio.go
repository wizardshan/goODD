package user

import (
	"github.com/tidwall/gjson"
)

type Bio struct {
	Value string
	Set   bool
	Mask  bool
}

func NewBio(v string) Bio {
	return Bio{Value: v, Set: true}
}

func (bio Bio) IsMask() bool {
	return bio.Mask
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
	results := gjson.GetManyBytes(data, "Value", "Mask")
	if results[0].Exists() {
		bio.Value = results[0].String()
		bio.Set = true
	}
	if results[1].Exists() {
		bio.Mask = results[1].Bool()
	}
	return nil
}
