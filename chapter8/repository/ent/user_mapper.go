package ent

import (
	"goODD/chapter8/domain"
)

func (entUser *User) Mapper() domain.User {
	var domUser domain.User
	if entUser == nil {
		return domUser
	}
	domUser.ID.SetTo(entUser.ID)
	domUser.Mobile.SetTo(entUser.Mobile)
	domUser.Age.SetTo(entUser.Age)
	domUser.Level.SetTo(entUser.Level)
	domUser.Nickname.SetTo(entUser.Nickname)
	domUser.CreateTime.SetTo(entUser.CreateTime)

	domUser.Set = true

	return domUser
}

func (entUsers Users) Mapper() domain.Users {
	size := len(entUsers)
	domUsers := make(domain.Users, size)
	for i := 0; i < size; i++ {
		domUsers[i] = entUsers[i].Mapper()
	}
	return domUsers
}