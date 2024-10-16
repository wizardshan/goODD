package domain

import (
	"goODD/chapter9/controller/response"
	"goODD/chapter9/domain/user"
	"goODD/chapter9/domain/vo"
	"time"
)

type Users []User

type User struct {
	ID         vo.ID
	Mobile     vo.Mobile
	Age        user.Age
	Level      user.Level
	Nickname   user.Nickname
	CreateTime vo.TimeRange

	Set bool
}

func (domUser User) IsZero() bool {
	return !domUser.Set
}

func (domUser User) Mapper() response.User {
	var respUser response.User
	domUser.ID.IsPresent(func(v int64) {
		respUser.ID = &v
	})
	domUser.Mobile.IsPresent(func(v string) {
		respUser.Mobile = &v
	})
	domUser.Age.IsPresent(func(v int64) {
		respUser.Age = &v
	})
	domUser.Level.IsPresent(func(v int64) {
		respUser.Level = &v
	})
	domUser.Nickname.IsPresent(func(v string) {
		respUser.Nickname = &v
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
