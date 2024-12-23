package domain

import (
	"goODD/chapter6/controller/response"
	"goODD/chapter6/domain/user"
	"goODD/chapter6/domain/vo"
	"time"
)

type Users []User

type User struct {
	ID         vo.ID
	Mobile     vo.Mobile
	Age        user.Age
	Level      user.Level
	Nickname   user.Nickname
	CreateTime time.Time
	Set        bool
}

func (domUser User) Mapper() response.User {
	var respUser response.User
	respUser.ID = domUser.ID.Value
	respUser.Mobile = domUser.Mobile.Value
	respUser.Age = domUser.Age.Value
	respUser.Level = domUser.Level.Value
	respUser.Nickname = domUser.Nickname.Value
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
