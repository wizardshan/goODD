package request

import (
	"goODD/chapter7/domain"
	"goODD/chapter7/domain/user"
)

type UserLogin struct {
	Mobile domain.Mobile `binding:"required"`
}

type UserOne struct {
	ID         domain.ID     `binding:"required"`
	HashID     user.HashID   `binding:"-"`
	Mobile     domain.Mobile `binding:"-"`
	Age        user.Age      `binding:"-"`
	Level      user.Level    `binding:"-"`
	Nickname   user.Nickname `binding:"-"`
	Avatar     user.Avatar   `binding:"-"`
	Bio        user.Bio      `binding:"-"`
	Amount     user.Amount   `binding:"-"`
	CreateTime domain.Time   `binding:"-"`
}

func (req UserOne) Mapper() domain.User {
	var domUser domain.User
	domUser.ID = req.ID
	domUser.HashID = req.HashID
	domUser.Mobile = req.Mobile
	domUser.Age = req.Age
	domUser.Level = req.Level
	domUser.Nickname = req.Nickname
	domUser.Avatar = req.Avatar
	domUser.Bio = req.Bio
	domUser.Amount = req.Amount
	domUser.CreateTime = req.CreateTime
	return domUser
}

type UserMany struct {
	ID         domain.ID     `binding:"-"`
	HashID     user.HashID   `binding:"-"`
	Mobile     domain.Mobile `binding:"-"`
	Age        user.Age      `binding:"-"`
	Level      user.Level    `binding:"-"`
	Nickname   user.Nickname `binding:"-"`
	Avatar     user.Avatar   `binding:"-"`
	Bio        user.Bio      `binding:"-"`
	Amount     user.Amount   `binding:"-"`
	CreateTime domain.Time   `binding:"-"`
}
