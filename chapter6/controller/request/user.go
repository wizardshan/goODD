package request

import (
	"errors"
	"goODD/chapter6/domain/user"
	"goODD/chapter6/domain/vo"
)

type UserLogin struct {
	Mobile vo.Mobile
}

func (req *UserLogin) Validate() error {
	return req.Mobile.Validate()
}

type UserOne struct {
	ID vo.ID
}

func (req *UserOne) Validate() error {
	return req.ID.Validate()
}

type UserMany struct {
	Mobile   vo.Mobile
	StartAge user.Age
	EndAge   user.Age
	Levels   user.Levels
	Nickname user.Nickname
}

func (req *UserMany) Validate() error {
	if err := req.Mobile.ValidateOmit(); err != nil {
		return err
	}
	if err := req.StartAge.ValidateOmit(); err != nil {
		return err
	}
	if err := req.EndAge.ValidateOmit(); err != nil {
		return err
	}
	for _, level := range req.Levels {
		if err := level.Validate(); err != nil {
			return err
		}
	}
	if err := req.Nickname.ValidateOmit(); err != nil {
		return err
	}
	if req.StartAge.Set && req.EndAge.Set && req.StartAge.Value >= req.EndAge.Value {
		return errors.New("年龄开始结束非法")
	}
	return nil
}
