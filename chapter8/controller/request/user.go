package request

import (
	"goODD/chapter8/domain/user"
	"goODD/chapter8/domain/vo"
)

type UserLogin struct {
	Mobile vo.Mobile
}

func (req UserLogin) Validate() error {
	if err := req.Mobile.Validate(); err != nil {
		return err
	}
	return nil
}

type UserOne struct {
	ID vo.ID
}

func (req UserOne) Validate() error {
	if err := req.ID.Validate(); err != nil {
		return err
	}
	return nil
}

type UserMany struct {
	Mobile   vo.Mobile
	Age      user.Age
	Level    user.Level
	Nickname user.Nickname
}

func (req UserMany) Validate() error {
	if err := req.Mobile.ValidateOmit(); err != nil {
		return err
	}
	if err := req.Age.RangeValidateOmit(); err != nil {
		return err
	}
	if err := req.Level.MultiValidateOmit(); err != nil {
		return err
	}
	if err := req.Nickname.FuzzyValidateOmit(); err != nil {
		return err
	}
	return nil
}
