package user

import (
	"github.com/tidwall/gjson"
	"goODD/chapter6/pkg/crypt"
)

type Password struct {
	Value     string `binding:"min=6,max=20"`
	Set       bool
	HashValue string
}

func (o *Password) Encode() {
	o.HashValue = crypt.PasswordHash(o.Value)
}

func (o *Password) Verify(v string) bool {
	return crypt.PasswordVerify(v, o.HashValue)
}

func (o *Password) IsPresent(f func(v string)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Password) SetTo(v string) {
	o.Set = true
	o.Value = v
}

func (o *Password) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.SetTo(gjson.ParseBytes(data).String())
		o.Encode()
		return nil
	}

	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		o.SetTo(result.String())
		o.Encode()
	}
	return nil
}
