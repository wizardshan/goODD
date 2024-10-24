package request

import "goODD/chapter1/controller/request/user"

type UserLoginV1 struct {
	Mobile  string `binding:"required,number,mobile"`
	Captcha string `binding:"required,number,len=4"`
}

type UserLoginV2 struct {
	MobileFieldV2
	CaptchaFieldV2
}

type UserLogin struct {
	MobileField
	CaptchaField
}

type UserOne struct {
	IDField
}

type UserMany struct {
	MobileField        `binding:"omitempty"`
	user.AgeField      `binding:"omitempty"`
	user.LevelField    `binding:"omitempty"`
	user.NicknameField `binding:"omitempty"`
}
