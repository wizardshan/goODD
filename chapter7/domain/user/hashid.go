package user

import (
	"errors"
	"github.com/tidwall/gjson"
	"goODD/chapter7/pkg/idx"
)

type HashID struct {
	Value    string
	ID       int64 `form:"-"`
	category int64
	Set      bool
	Mask     bool
}

func NewHashID(v string) HashID {
	return HashID{Value: v, Set: true}
}

func (hashID HashID) IsMask() bool {
	return hashID.Mask
}

func (hashID HashID) IsSet() bool {
	return hashID.Set
}

func (hashID HashID) IsPresent(f func(v string)) {
	if hashID.Set {
		f(hashID.Value)
	}
}

func (hashID HashID) Category() int64 {
	return idx.HashIDCategoryUser
}

func (hashID HashID) Encode() {
	hashID.Value = idx.Encode(hashID.ID, hashID.Category())
}

func (hashID HashID) Decode() (err error) {
	if hashID.ID, hashID.category, err = idx.Decode(hashID.Value); err != nil {
		return
	}
	if hashID.category != hashID.Category() {
		return errors.New("the id category error")
	}
	return
}

func (hashID *HashID) UnmarshalJSON(data []byte) error {
	results := gjson.GetManyBytes(data, "Value", "Mask")
	if results[0].Exists() {
		hashID.Value = results[0].String()
		if err := hashID.Decode(); err != nil {
			return err
		}
		hashID.Set = true
	}
	if results[1].Exists() {
		hashID.Mask = results[1].Bool()
	}
	return nil
}
