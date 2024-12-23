package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goODD/chapter6/controller/request"
	"goODD/chapter6/controller/response"
	"goODD/chapter6/domain"
	"goODD/chapter6/domain/vo"
	"goODD/chapter6/repository"
	"goODD/chapter6/repository/ent"
	"goODD/chapter6/repository/ent/user"
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
	idValue, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var id vo.ID
	id.SetTo(idValue)
	if err := id.Validate(); err != nil {
		return nil, err
	}

	domUser := ctr.repo.Find(c.Request.Context(), id)
	if !domUser.Set {
		return nil, errors.New("404 Not Found")
	}
	return domUser.Mapper(), nil
}

func (ctr *User) One(c *gin.Context) (response.Data, error) {
	req := new(request.UserOne)
	if err := ctr.Bind(c, req); err != nil {
		return nil, err
	}
	var qryUser domain.User
	qryUser.ID = req.ID

	domUser := ctr.repo.FindOne(c.Request.Context(), qryUser)
	if !domUser.Set {
		return nil, errors.New("404 Not Found")
	}
	return domUser.Mapper(), nil
}

func (ctr *User) Many(c *gin.Context) (response.Data, error) {
	req := new(request.UserMany)
	if err := ctr.Bind(c, req); err != nil {
		return nil, err
	}

	domUsers := ctr.repo.FetchMany(c.Request.Context(), func(opt *ent.UserQuery) {
		req.Mobile.IsPresent(func(v string) {
			opt.Where(user.Mobile(v))
		})
		req.StartAge.IsPresent(func(v int64) {
			opt.Where(user.AgeGTE(v))
		})
		req.EndAge.IsPresent(func(v int64) {
			opt.Where(user.AgeLTE(v))
		})
		req.Levels.IsPresent(func(v []int64) {
			opt.Where(user.LevelIn(v...))
		})
		req.Nickname.IsPresent(func(v string) {
			opt.Where(user.NicknameContains(v))
		})
	})
	return domUsers.Mapper(), nil
}
