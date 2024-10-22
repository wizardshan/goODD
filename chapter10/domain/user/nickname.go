package user

import (
	"chapter10/rpc/domain/user"
)

type Nickname struct {
	user.Nickname
	Set bool
}

func (o *Nickname) IsPresent(f func(v string)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Nickname) SetTo(v string) {
	o.Set = true
	o.Value = v
}

func (o *Nickname) SetToPb(v *user.Nickname) {
	if v != nil {
		o.SetTo(v.Value)
	}
}
