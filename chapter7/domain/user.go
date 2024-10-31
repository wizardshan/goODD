package domain

import (
	"goODD/chapter7/controller/response"
	"goODD/chapter7/domain/user"
	"goODD/chapter7/domain/vo"
	"time"
)

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
	Status     user.Status
	CreateTime time.Time
	UpdateTime time.Time

	Set bool
}

func (domUser *User) SmsRegister(mobile vo.Mobile) {
	domUser.Mobile = mobile
	domUser.Nickname.SetTo(mobile.GenNickname())
	domUser.Avatar.SetToDefault()
	domUser.CreateTime = time.Now()
}

func (domUser *User) Register(mobile vo.Mobile, password user.Password) {
	domUser.Mobile = mobile
	domUser.Nickname.SetTo(mobile.GenNickname())
	domUser.Avatar.SetToDefault()
	domUser.CreateTime = time.Now()
	domUser.Password = password
}

func (domUser User) Mapper() response.User {
	var respUser response.User
	if !domUser.Set {
		return respUser
	}
	respUser.ID = domUser.ID.Value
	respUser.Mobile = domUser.Mobile.Value
	respUser.HashID = domUser.HashID.Value
	respUser.Age = domUser.Age.Value
	respUser.Level = domUser.Level.Value
	respUser.LevelDesc = domUser.Level.Desc()
	respUser.Nickname = domUser.Nickname.Value
	respUser.Avatar = domUser.Avatar.Value
	respUser.Bio = domUser.Bio.Value
	respUser.Amount = domUser.Amount.ConvertYuan()
	respUser.CreateTime = domUser.CreateTime

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
