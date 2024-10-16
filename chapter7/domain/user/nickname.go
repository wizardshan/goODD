package user

import (
	"goODD/chapter7/domain/vo"
	"goODD/chapter7/pkg/validator"
)

type Nickname struct {
	vo.StringFuzzy
}

func NewNickname(v string) Nickname {
	var o Nickname
	o.SetTo(v)
	return o
}

func (o Nickname) validate(v string) error {
	return validator.Var(v, "min=2,max=10")
}

func (o Nickname) Validate() error {
	return o.validate(o.Value)
}

func (o Nickname) ValidateOmit() error {
	if o.Set {
		return o.validate(o.Value)
	}
	return nil
}

func (o Nickname) FuzzyValidate() error {
	return o.StringFuzzy.FuzzyValidate(o.validate)
}

func (o Nickname) FuzzyValidateOmit() error {
	return o.StringFuzzy.FuzzyValidateOmit(o.validate)
}
