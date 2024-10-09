package ent

import (
	"goODD/chapter7/domain"
	"goODD/chapter7/domain/user"
	entuser "goODD/chapter7/repository/ent/user"
)

func (entUser *User) MapperWithFields(fields domain.Fields) domain.User {
	var domUser domain.User
	if entUser == nil {
		return domUser
	}

	names, ok := fields.Names(entuser.Label)
	if !ok {
		return entUser.Mapper()
	}
	for _, name := range names {
		switch name {
		case "ID":
			domUser.ID = domain.NewID(entUser.ID)
		case "HashID":
			domUser.HashID = user.NewHashID(entUser.HashID)
		case "Mobile":
			domUser.Mobile = domain.NewMobile(entUser.Mobile)
		case "Password":
			domUser.Password = user.NewPassword(entUser.Password)
		case "Age":
			domUser.Age = user.NewAge(entUser.Age)
		case "Level":
			domUser.Level = user.NewLevel(entUser.Level)
		case "Nickname":
			domUser.Nickname = user.NewNickname(entUser.Nickname)
		case "Avatar":
			domUser.Avatar = user.NewAvatar(entUser.Avatar)
		case "Bio":
			domUser.Bio = user.NewBio(entUser.Bio)
		case "Amount":
			domUser.Amount = user.NewAmount(entUser.Amount)
		case "CreateTime":
			domUser.CreateTime = domain.NewTime(entUser.CreateTime)
		case "UpdateTime":
			domUser.UpdateTime = domain.NewTime(entUser.UpdateTime)
		}
	}
	return domUser
}

func (entUsers Users) MapperWithFields(fields domain.Fields) domain.Users {
	size := len(entUsers)
	domUsers := make(domain.Users, size)
	for i := 0; i < size; i++ {
		domUsers[i] = entUsers[i].MapperWithFields(fields)
	}
	return domUsers
}

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
