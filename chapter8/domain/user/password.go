package user

import (
	"chapter8/pkg/crypt"
	"chapter8/rpc/domain/user"
)

type Password struct {
	user.Password
	Set       bool
	HashValue string
}

func (o *Password) Encode() {
	o.HashValue = crypt.PasswordHash(o.Value)
}

func (o *Password) Verify(v string) bool {
	return crypt.PasswordVerify(v, o.HashValue)
}

func (o *Password) IsPresent(f func(v string)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Password) SetTo(v string) {
	o.Set = true
	o.Value = v
}

func (o *Password) SetToPb(v *user.Password) {
	if v != nil {
		o.SetTo(v.Value)
	}
}
