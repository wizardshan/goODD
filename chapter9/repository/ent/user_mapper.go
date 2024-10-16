package ent

import (
	"goODD/chapter9/domain"
	entuser "goODD/chapter9/repository/ent/user"
)

func (entUser *User) Mapper(fields domain.Fields) domain.User {
	var domUser domain.User
	if entUser == nil {
		return domUser
	}
	names, ok := fields.Names(entuser.Label)
	if !ok {
		names = []string{"ID", "Mobile", "Age", "Level", "Nickname", "CreateTime"}
	}

	for _, name := range names {
		switch name {
		case "ID":
			domUser.ID.SetTo(entUser.ID)
		case "Mobile":
			domUser.Mobile.SetTo(entUser.Mobile)
		case "Age":
			domUser.Age.SetTo(entUser.Age)
		case "Level":
			domUser.Level.SetTo(entUser.Level)
		case "Nickname":
			domUser.Nickname.SetTo(entUser.Nickname)
		case "CreateTime":
			domUser.CreateTime.SetTo(entUser.CreateTime)
		}
	}

	domUser.Set = true

	return domUser
}

func (entUsers Users) Mapper(fields domain.Fields) domain.Users {
	size := len(entUsers)
	domUsers := make(domain.Users, size)
	for i := 0; i < size; i++ {
		domUsers[i] = entUsers[i].Mapper(fields)
	}
	return domUsers
}
