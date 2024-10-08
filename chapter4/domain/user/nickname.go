package user

import (
	"github.com/tidwall/gjson"
)

type Nickname struct {
	Value string `binding:"min=2,max=10"`
	Set   bool
}

func (nickname *Nickname) UnmarshalJSON(data []byte) error {
	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		nickname.Value = result.String()
		nickname.Set = true
	}
	return nil
}
