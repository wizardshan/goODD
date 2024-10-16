package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goODD/chapter10/controller/request"
	"goODD/chapter10/controller/response"
	"goODD/chapter10/domain"
	"goODD/chapter10/repository"
	"goODD/chapter10/repository/ent"
	"goODD/chapter10/repository/ent/user"
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

func (ctr *User) Login(c *gin.Context) (response.Data, error) {
	req := new(request.UserLogin)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	mobileNotExist := ctr.repo.NotExist(c.Request.Context(), func(opt *ent.UserQuery) {
		opt.Where(user.Mobile(req.Mobile.Value))
	})
	if mobileNotExist {
		return nil, errors.New("手机号未注册")
	}

	domUser := ctr.repo.FetchOne(c.Request.Context(), func(opt *ent.UserQuery) {
		opt.Where(user.Mobile(req.Mobile.Value))
	})

	if !domUser.Password.Verify(req.Password.Value) {
		return nil, errors.New("密码错误")
	}

	return domUser.Mapper(), nil
}

func (ctr *User) SmsRegister(c *gin.Context) (response.Data, error) {
	req := new(request.UserSmsRegister)
	if err := ctr.Bind(c, req); err != nil {
		return nil, err
	}

	mobileExist := ctr.repo.Exist(c.Request.Context(), func(opt *ent.UserQuery) {
		opt.Where(user.Mobile(req.Mobile.Value))
	})
	if mobileExist {
		return nil, errors.New("手机号已注册")
	}

	cmdUser := domain.NewUser(
		req.Mobile,
	)

	domUser, err := ctr.repo.Register(c.Request.Context(), cmdUser)
	return domUser.Mapper(), err
}

func (ctr *User) Register(c *gin.Context) (response.Data, error) {
	req := new(request.UserRegister)
	if err := ctr.Bind(c, req); err != nil {
		return nil, err
	}

	mobileExist := ctr.repo.Exist(c.Request.Context(), func(opt *ent.UserQuery) {
		opt.Where(user.Mobile(req.Mobile.Value))
	})
	if mobileExist {
		return nil, errors.New("手机号已注册")
	}

	cmdUser := domain.NewUser(
		req.Mobile,
		domain.UserPassword(req.Password.Value),
	)

	domUser, err := ctr.repo.Register(c.Request.Context(), cmdUser)
	return domUser.Mapper(), err
}

func (ctr *User) Modify(c *gin.Context) (response.Data, error) {
	req := new(request.UserModify)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	var cmdUser domain.User
	cmdUser.ID.SetTo(1)
	cmdUser.Age = req.Age
	cmdUser.Nickname = req.Nickname
	cmdUser.Avatar = req.Avatar
	cmdUser.Bio = req.Bio
	ctr.repo.Modify(c.Request.Context(), cmdUser)
	return nil, nil
}

func (ctr *User) Cash(c *gin.Context) (response.Data, error) {
	req := new(request.UserCash)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	var userID int64 = 1
	amount := req.Amount
	domUser := ctr.repo.Fetch(c.Request.Context(), userID)
	if domUser.Amount.Insufficient(amount.Value) {
		return nil, errors.New("余额不足")
	}

	ctr.repo.Update(c.Request.Context(), func(opt *ent.UserUpdate) {
		opt.AddAmount(-amount.Value).Where(user.ID(userID))
	})

	return nil, nil
}
