package vo

import (
	"chapter8/rpc/domain"
)

type Captcha struct {
	domain.Captcha
	Set bool
}

func (o *Captcha) IsPresent(f func(v string)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Captcha) SetTo(v string) {
	o.Set = true
	o.Value = v
}

func (o *Captcha) SetToPb(v *domain.Captcha) {
	if v != nil {
		o.SetTo(v.Value)
	}
}
