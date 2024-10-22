package vo

import (
	"chapter10/rpc/domain"
)

type Mobile struct {
	domain.Mobile
	Set bool
}

func (o *Mobile) Mask() string {
	return o.Value[0:4] + "****" + o.Value[8:11]
}

func (o *Mobile) GenNickname() string {
	return o.Mask()
}

func (o *Mobile) IsPresent(f func(v string)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Mobile) SetTo(v string) {
	o.Set = true
	o.Value = v
}

func (o *Mobile) SetToPb(v *domain.Mobile) {
	if v != nil {
		o.SetTo(v.Value)
	}
}
