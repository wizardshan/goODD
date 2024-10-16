package user

import (
	"errors"
	"github.com/tidwall/gjson"
	"goODD/chapter10/domain/vo"
	"goODD/chapter10/pkg/crypt"
	"goODD/chapter10/pkg/validator"
)

type Password struct {
	vo.StringValue
	ReValue   string
	HashValue vo.StringValue
}

func (o Password) Encode() string {
	return crypt.PasswordHash(o.Value)
}

func (o Password) Verify(v string) bool {
	return crypt.PasswordVerify(v, o.HashValue.Value)
}

func (o Password) ValidateRepeat() error {
	if o.Value != o.ReValue {
		return errors.New("两次密码不一致")
	}
	return nil
}

func (o Password) validate(v string) error {
	return validator.Var(v, "min=6,max=20")
}

func (o Password) Validate() error {
	return o.validate(o.Value)
}

func (o Password) ValidateOmit() error {
	if o.Set {
		return o.validate(o.Value)
	}
	return nil
}

func (o *Password) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.SetTo(gjson.ParseBytes(data).String())
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "ReValue")
	if results[0].Exists() {
		o.SetTo(results[0].String())
	}
	if results[1].Exists() {
		o.ReValue = results[1].String()
	}
	return nil
}
