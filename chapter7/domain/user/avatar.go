package user

import (
	"chapter7/rpc/domain/user"
)

type Avatar struct {
	user.Avatar
	Set bool
}

func (o *Avatar) SetToDefault() {
	o.SetTo(o.Default())
}

func (o *Avatar) Default() string {
	return "avatar_default.png"
}

func (o *Avatar) IsPresent(f func(v string)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Avatar) SetTo(v string) {
	o.Set = true
	o.Value = v
}

func (o *Avatar) SetToPb(v *user.Avatar) {
	if v != nil {
		o.SetTo(v.Value)
	}
}
