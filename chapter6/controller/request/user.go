package request

import (
	"goODD/chapter6/domain"
	"goODD/chapter6/domain/user"
)

type UserLogin struct {
	Mobile domain.Mobile `binding:"required"`
}

type UserOne struct {
	ID domain.ID `binding:"required"`
}

type UserMany struct {
	Mobile   domain.Mobile `binding:"-"`
	Level    user.Level    `binding:"-"`
	Nickname user.Nickname `binding:"-"`
}
