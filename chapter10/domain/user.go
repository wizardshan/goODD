package domain

import (
	"goODD/chapter10/controller/response"
	"goODD/chapter10/domain/user"
	"goODD/chapter10/domain/vo"
)

type UserOption func(*User)

func UserNickname(v string) UserOption {
	return func(u *User) {
		u.Nickname.SetTo(v)
	}
}

func UserPassword(v string) UserOption {
	return func(u *User) {
		u.Password.SetTo(v)
	}
}

func NewUser(mobile vo.Mobile, options ...UserOption) User {
	var u User
	u.Mobile = mobile
	u.Nickname.SetTo(mobile.GenNickname())
	u.Avatar.SetToDefault()
	u.CreateTime.SetToNow()

	for _, option := range options {
		option(&u)
	}
	return u
}

type Users []User

type User struct {
	ID         vo.ID
	HashID     user.HashID
	Mobile     vo.Mobile
	Password   user.Password
	Age        user.Age
	Level      user.Level
	Nickname   user.Nickname
	Avatar     user.Avatar
	Bio        user.Bio
	Amount     user.Amount
	CreateTime vo.TimeRange
	UpdateTime vo.TimeValue

	Set bool
}

func (domUser User) IsZero() bool {
	return !domUser.Set
}

func (domUser User) Mapper() response.User {
	var respUser response.User
	respUser.ID = domUser.ID.Value
	respUser.Mobile = domUser.Mobile.Value
	respUser.HashID = domUser.HashID.Value
	respUser.Age = domUser.Age.Value
	respUser.Level = domUser.Level.Value
	respUser.LevelDesc = domUser.Level.Desc
	respUser.Nickname = domUser.Nickname.Value
	respUser.Avatar = domUser.Avatar.Value
	respUser.Bio = domUser.Bio.Value
	respUser.Amount = domUser.Amount.ConvertYuan()
	respUser.CreateTime = domUser.CreateTime.Value
	return respUser
}

func (domUsers Users) Mapper() response.Users {
	size := len(domUsers)
	respUsers := make(response.Users, size)
	for i := 0; i < size; i++ {
		respUsers[i] = domUsers[i].Mapper()
	}
	return respUsers
}
