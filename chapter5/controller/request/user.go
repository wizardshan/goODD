package request

import (
	"goODD/chapter5/domain/user"
	"goODD/chapter5/domain/vo"
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
	Age      user.Age
	Level    user.Level
	Nickname user.Nickname
}

func (req *UserMany) Validate() error {
	if err := req.Mobile.ValidateOmit(); err != nil {
		return err
	}
	if err := req.Age.ValidateOmit(); err != nil {
		return err
	}
	if err := req.Level.ValidateOmit(); err != nil {
		return err
	}
	if err := req.Nickname.ValidateOmit(); err != nil {
		return err
	}
	return nil
}
