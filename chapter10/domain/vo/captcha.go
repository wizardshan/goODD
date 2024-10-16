package vo

import "goODD/chapter10/pkg/validator"

type Captcha struct {
	StringValue
}

func (o Captcha) validate(v string) error {
	return validator.Var(v, "number,len=4")
}

func (o Captcha) Validate() error {
	return o.validate(o.Value)
}

func (o Captcha) ValidateOmit() error {
	if o.Set {
		return o.validate(o.Value)
	}
	return nil
}
