package user

import (
	"github.com/tidwall/gjson"
)

type Avatar struct {
	Value string
	Set   bool
	Mask  bool
}

func NewAvatar(v string) Avatar {
	return Avatar{Value: v, Set: true}
}

func (avatar Avatar) IsMask() bool {
	return avatar.Mask
}

func (avatar Avatar) IsSet() bool {
	return avatar.Set
}

func (avatar Avatar) IsPresent(f func(v string)) {
	if avatar.Set {
		f(avatar.Value)
	}
}

func (avatar *Avatar) UnmarshalJSON(data []byte) error {
	results := gjson.GetManyBytes(data, "Value", "Mask")
	if results[0].Exists() {
		avatar.Value = results[0].String()
		avatar.Set = true
	}
	if results[1].Exists() {
		avatar.Mask = results[1].Bool()
	}
	return nil
}
