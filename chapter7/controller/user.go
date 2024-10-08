package controller

import (
	"github.com/gin-gonic/gin"
	"goODD/chapter7/controller/request"
	"goODD/chapter7/controller/response"
	"goODD/chapter7/repository"
)

type User struct {
	repo *repository.User
}

func NewUser(repo *repository.User) *User {
	ctr := new(User)
	ctr.repo = repo
	return ctr
}

func (ctr *User) Login(c *gin.Context) (response.Data, error) {
	req := new(request.UserLogin)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (ctr *User) One(c *gin.Context) (response.Data, error) {
	req := new(request.UserOne)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	qryUser := req.Mapper()
	domUser := ctr.repo.Fetch(c.Request.Context(), qryUser)
	return domUser.Mapper(qryUser), nil
}

func (ctr *User) Many(c *gin.Context) (response.Data, error) {
	req := new(request.UserMany)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}
	return req, nil
}
