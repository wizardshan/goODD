package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goODD/chapter7/controller/request"
	"goODD/chapter7/controller/response"
	"goODD/chapter7/domain"
	"goODD/chapter7/domain/vo"
	"goODD/chapter7/repository"
	"goODD/chapter7/repository/ent"
	"goODD/chapter7/repository/ent/user"
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

func (ctr *User) Login(c *gin.Context) (response.Data, error) {
	req := new(request.UserLogin)
	if err := ctr.Bind(c, req); err != nil {
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

	var cmdUser domain.User
	cmdUser.SmsRegister(req.Mobile)
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

	var cmdUser domain.User
	cmdUser.Register(req.Mobile, req.Password)

	domUser, err := ctr.repo.Register(c.Request.Context(), cmdUser)
	return domUser.Mapper(), err
}

func (ctr *User) Modify(c *gin.Context) (response.Data, error) {
	req := new(request.UserModify)
	if err := ctr.Bind(c, req); err != nil {
		return nil, err
	}
	cmdUser := domain.User{
		ID:       req.ID,
		Age:      req.Age,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Bio:      req.Bio,
	}
	ctr.repo.Modify(c.Request.Context(), cmdUser)
	return nil, nil
}

func (ctr *User) Cash(c *gin.Context) (response.Data, error) {
	req := new(request.UserCash)
	if err := ctr.Bind(c, req); err != nil {
		return nil, err
	}

	var userID int64 = 1
	domUser := ctr.repo.Fetch(c.Request.Context(), userID)
	if domUser.Amount.Insufficient(req.Amount.Value) {
		return nil, errors.New("余额不足")
	}

	ctr.repo.Update(c.Request.Context(), func(opt *ent.UserUpdate) {
		opt.AddAmount(-req.Amount.Value).Where(user.ID(userID))
	})
	return nil, nil
}

func (ctr *User) ChangePassword(c *gin.Context) (response.Data, error) {
	req := new(request.UserChangePassword)
	if err := ctr.Bind(c, req); err != nil {
		return nil, err
	}
	cmdUser := domain.User{
		Password: req.Password,
	}
	var userID int64 = 1
	cmdUser.ID.SetTo(userID)
	ctr.repo.Modify(c.Request.Context(), cmdUser)
	return nil, nil
}
