package ent

import (
	"goODD/chapter7/domain"
	"goODD/chapter7/domain/user"
)

func (entUser *User) Mapper() domain.User {
	var domUser domain.User
	if entUser == nil {
		return domUser
	}
	domUser.ID = domain.NewID(entUser.ID)
	domUser.HashID = user.NewHashID(entUser.HashID)
	domUser.Mobile = domain.NewMobile(entUser.Mobile)
	domUser.Password = user.NewPassword(entUser.Password)
	domUser.Age = user.NewAge(entUser.Age)
	domUser.Level = user.NewLevel(entUser.Level)
	domUser.Nickname = user.NewNickname(entUser.Nickname)
	domUser.Avatar = user.NewAvatar(entUser.Avatar)
	domUser.Bio = user.NewBio(entUser.Bio)
	domUser.Amount = user.NewAmount(entUser.Amount)
	domUser.CreateTime = domain.NewTime(entUser.CreateTime)
	domUser.UpdateTime = domain.NewTime(entUser.UpdateTime)
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
