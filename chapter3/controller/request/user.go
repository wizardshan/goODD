package request

import (
	"goODD/chapter3/domain"
	"goODD/chapter3/domain/user"
)

type UserLogin struct {
	Mobile domain.Mobile `binding:"required"`
}

type UserOne struct {
	ID domain.ID `binding:"required"`
}

type UserMany struct {
	Mobile   domain.Mobile `binding:"omitempty"`
	Level    user.Level    `binding:"omitempty"`
	Nickname user.Nickname `binding:"omitempty"`
}
