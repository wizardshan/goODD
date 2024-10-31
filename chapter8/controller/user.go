package controller

import (
	"chapter8/domain"
	"chapter8/domain/user"
	"chapter8/domain/vo"
	"chapter8/repository"
	"chapter8/repository/ent"
	entUser "chapter8/repository/ent/user"
	"chapter8/rpc/request"
	"chapter8/rpc/response"
	"chapter8/rpc/rpcconnect"
	"connectrpc.com/connect"
	"context"
	"errors"
)

type User struct {
	repo *repository.User
	rpcconnect.UnimplementedUserHandler
}

func NewUser(repo *repository.User) *User {
	ctr := new(User)
	ctr.repo = repo
	return ctr
}

func (ctr *User) One(ctx context.Context, req *connect.Request[request.UserOne]) (*connect.Response[response.UserOne], error) {
	domUser := ctr.repo.Fetch(ctx, req.Msg.ID.GetValue())
	if !domUser.Set {
		return nil, errors.New("404 Not Found")
	}

	return connect.NewResponse(&response.UserOne{User: domUser.Mapper()}), nil
}

func (ctr *User) Many(ctx context.Context, req *connect.Request[request.UserMany]) (*connect.Response[response.UserMany], error) {
	var qryMobile vo.Mobile
	qryMobile.SetToPb(req.Msg.Mobile)
	var qryStartAge user.Age
	qryStartAge.SetToPb(req.Msg.StartAge)
	var qryEndAge user.Age
	qryEndAge.SetToPb(req.Msg.EndAge)
	var qryLevels user.Levels = req.Msg.Levels
	var qryNickname user.Nickname
	qryNickname.SetToPb(req.Msg.Nickname)

	domUsers := ctr.repo.FetchMany(ctx, func(opt *ent.UserQuery) {
		qryMobile.IsPresent(func(v string) {
			opt.Where(entUser.Mobile(v))
		})
		qryStartAge.IsPresent(func(v int64) {
			opt.Where(entUser.AgeGTE(v))
		})
		qryEndAge.IsPresent(func(v int64) {
			opt.Where(entUser.AgeLTE(v))
		})
		qryLevels.IsPresent(func(v []int64) {
			opt.Where(entUser.LevelIn(v...))
		})
		qryNickname.IsPresent(func(v string) {
			opt.Where(entUser.NicknameContains(v))
		})
	})

	return connect.NewResponse(&response.UserMany{List: domUsers.Mapper()}), nil
}

func (ctr *User) Login(ctx context.Context, req *connect.Request[request.UserLogin]) (*connect.Response[response.UserLogin], error) {

	var qryMobile vo.Mobile
	qryMobile.SetToPb(req.Msg.Mobile)

	mobileNotExist := ctr.repo.NotExist(ctx, func(opt *ent.UserQuery) {
		opt.Where(entUser.Mobile(qryMobile.Value))
	})
	if mobileNotExist {
		return nil, errors.New("手机号未注册")
	}

	domUser := ctr.repo.FetchOne(ctx, func(opt *ent.UserQuery) {
		opt.Where(entUser.Mobile(qryMobile.Value))
	})

	if !domUser.Password.Verify(req.Msg.Password.GetValue()) {
		return nil, errors.New("密码错误")
	}

	return connect.NewResponse(&response.UserLogin{User: domUser.Mapper()}), nil
}

func (ctr *User) SmsRegister(ctx context.Context, req *connect.Request[request.UserSmsRegister]) (*connect.Response[response.UserSmsRegister], error) {
	var qryMobile vo.Mobile
	qryMobile.SetToPb(req.Msg.Mobile)
	mobileExist := ctr.repo.Exist(ctx, func(opt *ent.UserQuery) {
		opt.Where(entUser.Mobile(qryMobile.Value))
	})
	if mobileExist {
		return nil, errors.New("手机号已注册")
	}

	cmdUser := domain.NewUser(
		qryMobile,
	)

	domUser, err := ctr.repo.Register(ctx, cmdUser)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&response.UserSmsRegister{User: domUser.Mapper()}), nil
}

func (ctr *User) Register(ctx context.Context, req *connect.Request[request.UserRegister]) (*connect.Response[response.UserRegister], error) {
	var qryMobile vo.Mobile
	qryMobile.SetToPb(req.Msg.Mobile)
	var qryPassword user.Password
	qryPassword.SetToPb(req.Msg.Password)

	mobileExist := ctr.repo.Exist(ctx, func(opt *ent.UserQuery) {
		opt.Where(entUser.Mobile(qryMobile.Value))
	})
	if mobileExist {
		return nil, errors.New("手机号已注册")
	}

	cmdUser := domain.NewUser(
		qryMobile,
		domain.UserPassword(qryPassword),
	)

	domUser, err := ctr.repo.Register(ctx, cmdUser)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&response.UserRegister{User: domUser.Mapper()}), nil
}

func (ctr *User) Modify(ctx context.Context, req *connect.Request[request.UserModify]) (*connect.Response[response.UserModify], error) {
	var cmdUser domain.User
	cmdUser.ID.SetToPb(req.Msg.ID)
	cmdUser.Age.SetToPb(req.Msg.Age)
	cmdUser.Nickname.SetToPb(req.Msg.Nickname)
	cmdUser.Avatar.SetToPb(req.Msg.Avatar)
	cmdUser.Bio.SetToPb(req.Msg.Bio)
	ctr.repo.Modify(ctx, cmdUser)

	return connect.NewResponse(&response.UserModify{}), nil
}

func (ctr *User) Cash(ctx context.Context, req *connect.Request[request.UserCash]) (*connect.Response[response.UserCash], error) {
	var qryAmount user.Amount
	qryAmount.SetToPb(req.Msg.Amount)
	if err := qryAmount.CashLimit(); err != nil {
		return nil, err
	}

	var userID int64 = 1
	domUser := ctr.repo.Fetch(ctx, userID)
	if domUser.Amount.Insufficient(qryAmount.Value) {
		return nil, errors.New("余额不足")
	}

	ctr.repo.Update(ctx, func(opt *ent.UserUpdate) {
		opt.AddAmount(-qryAmount.Value).Where(entUser.ID(userID))
	})
	return connect.NewResponse(&response.UserCash{}), nil
}
