package ent

import (
	"chapter10/domain"
)

func (entUser *User) Mapper() domain.User {
	var domUser domain.User
	if entUser == nil {
		return domUser
	}
	domUser.ID.SetTo(entUser.ID)
	domUser.HashID.SetTo(entUser.HashID)
	domUser.Mobile.SetTo(entUser.Mobile)
	domUser.Password.HashValue = entUser.Password
	domUser.Age.SetTo(entUser.Age)
	domUser.Level.SetTo(entUser.Level)
	domUser.Nickname.SetTo(entUser.Nickname)
	domUser.Avatar.SetTo(entUser.Avatar)
	domUser.Bio.SetTo(entUser.Bio)
	domUser.Amount.SetTo(entUser.Amount)
	domUser.CreateTime = entUser.CreateTime
	domUser.UpdateTime = entUser.UpdateTime

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
