package vo

import (
	"chapter7/rpc/domain"
)

type ID struct {
	domain.ID
	Set bool
}

func (o *ID) IsPresent(f func(v int64)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *ID) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

func (o *ID) SetToPb(v *domain.ID) {
	if v != nil {
		o.SetTo(v.Value)
	}
}
