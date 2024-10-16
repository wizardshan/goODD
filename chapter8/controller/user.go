package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goODD/chapter8/controller/request"
	"goODD/chapter8/controller/response"
	"goODD/chapter8/domain"
	"goODD/chapter8/repository"
	"goODD/chapter8/repository/ent"
	"goODD/chapter8/repository/ent/user"
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

	domUser := ctr.repo.Fetch(c.Request.Context(), id)
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

	domUser := ctr.repo.Fetch(c.Request.Context(), req.ID.Value)
	if domUser.IsZero() {
		return nil, errors.New("404 Not Found")
	}

	return domUser.Mapper(), nil
}

func (ctr *User) Many1(c *gin.Context) (response.Data, error) {
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
	domUsers := ctr.repo.FindMany(c.Request.Context(), qryUser)
	return domUsers.Mapper(), nil
}

func (ctr *User) Many2(c *gin.Context) (response.Data, error) {
	req := new(request.UserMany)
	if err := ctr.Bind(c, req); err != nil {
		return nil, err
	}

	domUsers := ctr.repo.FetchMany(c.Request.Context(), func(opt *ent.UserQuery) {
		opt.Where(
			user.Mobile(req.Mobile.Value),
			user.AgeGTE(req.Age.Start.Value),
			user.AgeLTE(req.Age.End.Value),
			user.LevelIn(req.Level.Values...),
			user.NicknameContains(req.Nickname.Keyword.Value),
		)
	})
	return domUsers.Mapper(), nil
}

func (ctr *User) Many3(c *gin.Context) (response.Data, error) {
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
		Query: qryUser,
		Option: func(opt *ent.UserQuery) {
			opt.Where(user.Nickname("1300****001"))
		},
	})
	return domUsers.Mapper(), nil
}
