package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goODD/chapter9/controller/request"
	"goODD/chapter9/controller/response"
	"goODD/chapter9/domain"
	"goODD/chapter9/repository"
	"goODD/chapter9/repository/ent"
	"goODD/chapter9/repository/ent/user"
	"strconv"
)

type User struct {
	ctr
	repo *repository.User
}

func NewUser(repo *repository.User) *User {
	ctr := new(User)
	ctr.repo = repo
	return ctr
}

func (ctr *User) FetchOne(c *gin.Context) (response.Data, error) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	req := new(request.UserOne)
	req.ID.SetTo(id)
	if err := ctr.Bind(c, req); err != nil {
		return nil, err
	}

	domUser := ctr.repo.One(c.Request.Context(), repository.UserWrapper{
		Option: func(opt *ent.UserQuery) {
			opt.Where(user.ID(id))
		},
		Fields: c.QueryMap("Fields"),
	})
	if domUser.IsZero() {
		return nil, errors.New("404 Not Found")
	}
	return domUser.Mapper(), nil
}

func (ctr *User) One(c *gin.Context) (response.Data, error) {
	req := new(request.UserOne)
	if err := ctr.Bind(c, req); err != nil {
		return nil, err
	}

	domUser := ctr.repo.One(c.Request.Context(), repository.UserWrapper{
		Option: func(opt *ent.UserQuery) {
			opt.Where(user.ID(req.ID.Value))
		},
		Fields: c.QueryMap("Fields"),
	})
	if domUser.IsZero() {
		return nil, errors.New("404 Not Found")
	}
	return domUser.Mapper(), nil
}

func (ctr *User) Many(c *gin.Context) (response.Data, error) {
	req := new(request.UserMany)
	if err := ctr.Bind(c, req); err != nil {
		return nil, err
	}

	qryUser := domain.User{
		Mobile:   req.Mobile,
		Age:      req.Age,
		Level:    req.Level,
		Nickname: req.Nickname,
	}

	domUsers := ctr.repo.Many(c.Request.Context(), repository.UserWrapper{
		Query:  qryUser,
		Fields: c.QueryMap("Fields"),
	})
	return domUsers.Mapper(), nil
}
