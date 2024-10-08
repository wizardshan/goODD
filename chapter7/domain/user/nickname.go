package user

import (
	"github.com/tidwall/gjson"
)

type Nickname struct {
	Value string `binding:"min=2,max=10"`
	Set   bool
	Mask  bool
}

func NewNickname(v string) Nickname {
	return Nickname{Value: v, Set: true}
}

func (nickname Nickname) IsMask() bool {
	return nickname.Mask
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
	results := gjson.GetManyBytes(data, "Value", "Mask")
	if results[0].Exists() {
		nickname.Value = results[0].String()
		nickname.Set = true
	}
	if results[1].Exists() {
		nickname.Mask = results[1].Bool()
	}
	return nil
}
