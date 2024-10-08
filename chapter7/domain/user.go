package domain

import (
	"goODD/chapter7/controller/response"
	"goODD/chapter7/domain/user"
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

func (domUser User) IsMask() bool {
	return domUser.ID.Mask ||
		domUser.Mobile.Mask ||
		domUser.HashID.Mask ||
		domUser.Age.Mask ||
		domUser.Level.Mask ||
		domUser.Nickname.Mask ||
		domUser.Avatar.Mask ||
		domUser.Bio.Mask ||
		domUser.Amount.Mask ||
		domUser.CreateTime.Mask
}

func (domUser User) IsZero() bool {
	var zero User
	return zero == domUser
}

func (domUser User) Mapper(qryUser User) response.User {
	var respUser response.User
	if domUser.IsZero() {
		return respUser // or `err: 404 not found`
	}

	if !qryUser.IsMask() || qryUser.ID.IsMask() {
		respUser.ID = &domUser.ID.Value
	}
	if !qryUser.IsMask() || qryUser.Mobile.IsMask() {
		respUser.Mobile = &domUser.Mobile.Value
	}
	if !qryUser.IsMask() || qryUser.HashID.IsMask() {
		respUser.HashID = &domUser.HashID.Value
	}
	if !qryUser.IsMask() || qryUser.Age.IsMask() {
		respUser.Age = &domUser.Age.Value
	}
	if !qryUser.IsMask() || qryUser.Level.IsMask() {
		respUser.Level = &domUser.Level.Value
		respUser.LevelDesc = &domUser.Level.Desc
	}
	if !qryUser.IsMask() || qryUser.Nickname.IsMask() {
		respUser.Nickname = &domUser.Nickname.Value
	}
	if !qryUser.IsMask() || qryUser.Avatar.IsMask() {
		respUser.Avatar = &domUser.Avatar.Value
	}
	if !qryUser.IsMask() || qryUser.Bio.IsMask() {
		respUser.Bio = &domUser.Bio.Value
	}
	if !qryUser.IsMask() || qryUser.Amount.IsMask() {
		respUser.Amount = &domUser.Amount.Value
	}
	if !qryUser.IsMask() || qryUser.CreateTime.IsMask() {
		respUser.CreateTime = &domUser.CreateTime.Value
	}
	return respUser
}

func (domUsers Users) Mapper(queryUser User) response.Users {
	size := len(domUsers)
	respUsers := make(response.Users, size)
	for i := 0; i < size; i++ {
		respUsers[i] = domUsers[i].Mapper(queryUser)
	}
	return respUsers
}
