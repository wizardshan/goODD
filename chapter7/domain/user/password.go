package user

import "github.com/tidwall/gjson"

type Password struct {
	Value     string `binding:"min=6"`
	ReValue   string `binding:"eqfield=Value"`
	HashValue string
	Set       bool
}

func NewPassword(v string) Password {
	return Password{Value: v, Set: true}
}

func (password Password) IsSet() bool {
	return password.Set
}

func (password Password) IsPresent(f func(v string)) {
	if password.Set {
		f(password.Value)
	}
}

func (password *Password) UnmarshalJSON(data []byte) error {
	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		password.Value = result.String()
		password.Set = true
	}
	return nil
}
