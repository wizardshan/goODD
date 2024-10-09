package user

import (
	"github.com/tidwall/gjson"
)

type Nickname struct {
	Value string `binding:"min=2,max=10"`
	Set   bool
}

func NewNickname(v string) Nickname {
	return Nickname{Value: v, Set: true}
}

func (nickname Nickname) IsSet() bool {
	return nickname.Set
}

func (nickname Nickname) IsPresent(f func(v string)) {
	if nickname.Set {
		f(nickname.Value)
	}
}

func (nickname *Nickname) UnmarshalJSON(data []byte) error {
	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		nickname.Value = result.String()
		nickname.Set = true
	}
	return nil
}
