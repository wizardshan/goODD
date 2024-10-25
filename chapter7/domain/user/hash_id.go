package user

import (
	"chapter7/pkg/idx"
	"chapter7/rpc/domain/user"
	"errors"
)

type HashID struct {
	user.HashID
	Set      bool
	ID       int64
	category int64
}

func (o *HashID) Category() int64 {
	return idx.HashIDCategoryUser
}

func (o *HashID) Encode() {
	o.Value = idx.Encode(o.ID, o.Category())
}

func (o *HashID) Decode() (err error) {
	if o.ID, o.category, err = idx.Decode(o.Value); err != nil {
		return
	}
	if o.category != o.Category() {
		return errors.New("the id category error")
	}
	return
}

func (o *HashID) IsPresent(f func(v string)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *HashID) SetTo(v string) {
	o.Set = true
	o.Value = v
}

func (o *HashID) SetToPb(v *user.HashID) {
	if v != nil {
		o.SetTo(v.Value)
	}
}
