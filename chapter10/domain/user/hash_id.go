package user

import (
	"errors"
	"goODD/chapter10/domain/vo"
	"goODD/chapter10/pkg/idx"
	"goODD/chapter10/pkg/validator"
)

type HashID struct {
	vo.StringValue
	ID       int64
	category int64
}

func (o HashID) Category() int64 {
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

func (o HashID) validate(v string) error {
	return validator.Var(v, "max=200")
}

func (o HashID) Validate() error {
	return o.validate(o.Value)
}

func (o HashID) ValidateOmit() error {
	if o.Set {
		return o.validate(o.Value)
	}
	return nil
}
