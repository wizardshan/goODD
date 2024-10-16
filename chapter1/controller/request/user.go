package request

import "goODD/chapter1/controller/request/user"

type UserLoginV1 struct {
	Mobile string `binding:"required,number,mobile"`
}

type UserLoginV2 struct {
	MobileFieldV2
}

type UserLogin struct {
	MobileField `binding:"required"`
}

type UserOne struct {
	IDField `binding:"required"`
}

type UserMany struct {
	MobileField        `binding:"omitempty"`
	user.AgeField      `binding:"omitempty"`
	user.LevelField    `binding:"omitempty"`
	user.NicknameField `binding:"omitempty"`
}
