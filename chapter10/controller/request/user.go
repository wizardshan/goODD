package request

import (
	"goODD/chapter10/domain/user"
	"goODD/chapter10/domain/vo"
)

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

type UserLogin struct {
	Mobile   vo.Mobile
	Password user.Password
}

func (req UserLogin) Validate() error {
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

func (req UserSmsRegister) Validate() error {
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

func (req UserRegister) Validate() error {
	if err := req.Mobile.Validate(); err != nil {
		return err
	}
	if err := req.Password.Validate(); err != nil {
		return err
	}
	if err := req.Password.ValidateRepeat(); err != nil {
		return err
	}
	return nil
}

type UserModify struct {
	Age      user.Age
	Nickname user.Nickname
	Avatar   user.Avatar
	Bio      user.Bio
}

func (req UserModify) Validate() error {
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

type UserCash struct {
	Amount user.Amount
}

func (req UserCash) Validate() error {
	if err := req.Amount.CashLimit(); err != nil {
		return err
	}
	return nil
}
