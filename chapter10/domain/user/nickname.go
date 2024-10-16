package user

import (
	"goODD/chapter10/domain/vo"
	"goODD/chapter10/pkg/validator"
)

type Nickname struct {
	vo.StringFuzzy
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
