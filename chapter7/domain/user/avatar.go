package user

import (
	"github.com/tidwall/gjson"
)

type Avatar struct {
	Value string
	Set   bool
}

func NewAvatar(v string) Avatar {
	return Avatar{Value: v, Set: true}
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
	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		avatar.Value = result.String()
		avatar.Set = true
	}
	return nil
}
