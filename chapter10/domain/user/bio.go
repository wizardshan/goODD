package user

import "chapter10/rpc/domain/user"

type Bio struct {
	user.Bio
	Set bool
}

func (o *Bio) IsPresent(f func(v string)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Bio) SetTo(v string) {
	o.Set = true
	o.Value = v
}

func (o *Bio) SetToPb(v *user.Bio) {
	if v != nil {
		o.SetTo(v.Value)
	}
}
