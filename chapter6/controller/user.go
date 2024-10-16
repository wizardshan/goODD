package controller

import (
	"github.com/gin-gonic/gin"
	"goODD/chapter6/controller/request"
	"goODD/chapter6/controller/response"
)

type User struct {
	ctr
}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Login(c *gin.Context) (response.Data, error) {
	req := new(request.UserLogin)
	if err := ctr.Bind(c, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (ctr *User) One(c *gin.Context) (response.Data, error) {
	req := new(request.UserOne)
	if err := ctr.Bind(c, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (ctr *User) Many(c *gin.Context) (response.Data, error) {
	req := new(request.UserMany)
	if err := ctr.Bind(c, req); err != nil {
		return nil, err
	}
	return req, nil
}
