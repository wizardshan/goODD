package user

import (
	"github.com/tidwall/gjson"
	"goODD/chapter7/domain/vo"
	"goODD/chapter7/pkg/crypt"
	"goODD/chapter7/pkg/validate"
)

type Password struct {
	vo.StringValue
	Hash vo.StringValue
}

func (o *Password) Encode() {
	o.Hash.SetTo(crypt.PasswordHash(o.Value))
}

func (o *Password) Verify(v string) bool {
	return crypt.PasswordVerify(v, o.Hash.Value)
}

func (o *Password) ValidateOmit() error {
	if o.Set {
		return o.Validate()
	}
	return nil
}

func (o *Password) Validate() error {
	return validate.Var(o.Value, "min=6,max=20")
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
