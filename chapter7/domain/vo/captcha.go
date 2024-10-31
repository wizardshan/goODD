package vo

import "goODD/chapter7/pkg/validate"

type Captcha struct {
	Value string `binding:"number,len=4"`
	Set   bool
}

func (o *Captcha) ValidateOmit() error {
	if o.Set {
		return o.Validate()
	}
	return nil
}

func (o *Captcha) Validate() error {
	return validate.Var(o.Value, "number,len=4")
}
