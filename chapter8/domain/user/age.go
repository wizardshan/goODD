package user

import (
	"chapter8/rpc/domain/user"
)

type Age struct {
	user.Age
	Set bool
}

func (o *Age) IsPresent(f func(v int64)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Age) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

func (o *Age) SetToPb(v *user.Age) {
	if v != nil {
		o.SetTo(v.Value)
	}
}
