package user

import (
	"errors"
	"github.com/tidwall/gjson"
	"goODD/chapter6/pkg/idx"
)

type HashID struct {
	Value    string `binding:"min=2,max=10"`
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

func (o *HashID) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.SetTo(gjson.ParseBytes(data).String())
		return o.Decode()
	}

	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		o.SetTo(result.String())
		return o.Decode()
	}
	return nil
}
