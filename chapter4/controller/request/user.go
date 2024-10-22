package request

import (
	"goODD/chapter4/domain/user"
	"goODD/chapter4/domain/vo"
)

type UserLogin struct {
	Mobile vo.Mobile
}

type UserOne struct {
	ID vo.ID
}

type UserMany struct {
	Mobile   vo.Mobile     `binding:"omitempty"`
	Age      user.Age      `binding:"omitempty"`
	Level    user.Level    `binding:"omitempty"`
	Nickname user.Nickname `binding:"omitempty"`
}
