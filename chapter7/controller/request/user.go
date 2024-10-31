package request

import (
	"errors"
	"goODD/chapter7/domain/user"
	"goODD/chapter7/domain/vo"
)

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

type UserLogin struct {
	Mobile   vo.Mobile
	Password user.Password
}

func (req *UserLogin) Validate() error {
	if err := req.Mobile.Validate(); err != nil {
		return err
	}
	if err := req.Password.Validate(); err != nil {
		return err
	}
	return nil
}

type UserSmsRegister struct {
	Mobile  vo.Mobile
	Captcha vo.Captcha
}

func (req *UserSmsRegister) Validate() error {
	if err := req.Mobile.Validate(); err != nil {
		return err
	}
	if err := req.Captcha.Validate(); err != nil {
		return err
	}
	return nil
}

type UserRegister struct {
	Mobile   vo.Mobile
	Password user.Password
}

func (req *UserRegister) Validate() error {
	if err := req.Mobile.Validate(); err != nil {
		return err
	}
	if err := req.Password.Validate(); err != nil {
		return err
	}
	return nil
}

type UserModify struct {
	ID       vo.ID
	Age      user.Age
	Nickname user.Nickname
	Avatar   user.Avatar
	Bio      user.Bio
}

func (req *UserModify) Validate() error {
	if err := req.ID.Validate(); err != nil {
		return err
	}
	if err := req.Age.Validate(); err != nil {
		return err
	}
	if err := req.Nickname.Validate(); err != nil {
		return err
	}
	if err := req.Avatar.Validate(); err != nil {
		return err
	}
	if err := req.Bio.Validate(); err != nil {
		return err
	}
	return nil
}

type UserChangePassword struct {
	Password user.Password
}

func (req *UserChangePassword) Validate() error {
	return req.Password.Validate()
}

type UserCash struct {
	Amount user.Amount
}

func (req *UserCash) Validate() error {
	return req.Amount.CashLimit()
}
