package request

import (
	"goODD/chapter6/domain/user"
	"goODD/chapter6/domain/vo"
)

type UserOne struct {
	ID vo.ID
}

type UserMany struct {
	Mobile   vo.Mobile     `binding:"omitempty"`
	StartAge user.Age      `binding:"omitempty"`
	EndAge   user.Age      `binding:"omitempty,gtfield=Start"` // 不起作用，struct类型默认只支持time类型 需要修改源码 baked_in.go isGtField
	Levels   user.Levels   `binding:"omitempty,dive"`
	Nickname user.Nickname `binding:"omitempty"`
}

type UserLogin struct {
	Mobile   vo.Mobile
	Password user.Password
}

type UserSmsRegister struct {
	Mobile  vo.Mobile
	Captcha vo.Captcha
}

type UserRegister struct {
	Mobile   vo.Mobile
	Password user.Password
}

type UserModify struct {
	ID       vo.ID
	Age      user.Age
	Nickname user.Nickname
	Avatar   user.Avatar
	Bio      user.Bio
}

type UserCash struct {
	Amount user.Amount
}
