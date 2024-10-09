package domain

import (
	"goODD/chapter7/controller/response"
	"goODD/chapter7/domain/user"
	"time"
)

type Users []User

type User struct {
	ID         ID
	HashID     user.HashID
	Mobile     Mobile
	Password   user.Password
	Age        user.Age
	Level      user.Level
	Nickname   user.Nickname
	Avatar     user.Avatar
	Bio        user.Bio
	Amount     user.Amount
	CreateTime Time
	UpdateTime Time
}

func (domUser User) IsZero() bool {
	var zero User
	return zero == domUser
}

func (domUser User) Mapper() response.User {
	var respUser response.User
	if domUser.IsZero() {
		return respUser // or `err: 404 not found`
	}

	domUser.ID.IsPresent(func(v int64) {
		respUser.ID = &v
	})
	domUser.Mobile.IsPresent(func(v string) {
		respUser.Mobile = &v
	})
	domUser.HashID.IsPresent(func(v string) {
		respUser.HashID = &v
	})
	domUser.Age.IsPresent(func(v int64) {
		respUser.Age = &v
	})
	domUser.Level.IsPresent(func(v int64, desc string) {
		respUser.Level = &v
		respUser.LevelDesc = &desc
	})
	domUser.Nickname.IsPresent(func(v string) {
		respUser.Nickname = &v
	})
	domUser.Avatar.IsPresent(func(v string) {
		respUser.Avatar = &v
	})
	domUser.Bio.IsPresent(func(v string) {
		respUser.Bio = &v
	})
	domUser.Amount.IsPresent(func(v int64) {
		respUser.Amount = &v
	})
	domUser.CreateTime.IsPresent(func(v time.Time) {
		respUser.CreateTime = &v
	})
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
