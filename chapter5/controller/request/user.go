package request

import (
	"goODD/chapter5/domain/user"
	"goODD/chapter5/domain/vo"
)

type UserLogin struct {
	Mobile vo.Mobile
}

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
