package repository

import (
	"context"
	"goODD/chapter10/domain"
	"goODD/chapter10/repository/ent"
	"goODD/chapter10/repository/ent/user"
	"time"
)

type UserWrapper struct {
	Query  domain.User
	Option func(opt *ent.UserQuery)
}

func (repo *User) Fetch(ctx context.Context, id int64) domain.User {
	return repo.fetch(ctx, repo.db, id).Mapper()
}

func (repo *User) fetch(ctx context.Context, db *ent.Client, id int64) *ent.User {
	return repo.fetchOne(ctx, db, func(opt *ent.UserQuery) {
		opt.Where(user.ID(id))
	})
}

func (repo *User) FetchIn(ctx context.Context, ids ...int64) domain.User {
	return repo.fetchIn(ctx, repo.db, ids...).Mapper()
}

func (repo *User) fetchIn(ctx context.Context, db *ent.Client, ids ...int64) *ent.User {
	return repo.fetchOne(ctx, db, func(opt *ent.UserQuery) {
		opt.Where(user.IDIn(ids...))
	})
}

func (repo *User) FetchOne(ctx context.Context, option func(opt *ent.UserQuery)) domain.User {
	return repo.fetchOne(ctx, repo.db, option).Mapper()
}

func (repo *User) fetchOne(ctx context.Context, db *ent.Client, option func(opt *ent.UserQuery)) *ent.User {
	var wrapper UserWrapper
	wrapper.Option = option
	return repo.one(ctx, db, wrapper)
}

func (repo *User) FetchMany(ctx context.Context, option func(opt *ent.UserQuery)) domain.Users {
	return repo.fetchMany(ctx, repo.db, option).Mapper()
}

func (repo *User) fetchMany(ctx context.Context, db *ent.Client, option func(opt *ent.UserQuery)) ent.Users {
	var wrapper UserWrapper
	wrapper.Option = option
	return repo.many(ctx, db, wrapper)
}

func (repo *User) FindOne(ctx context.Context, qryUser domain.User) domain.User {
	return repo.findOne(ctx, repo.db, qryUser).Mapper()
}

func (repo *User) findOne(ctx context.Context, db *ent.Client, qryUser domain.User) *ent.User {
	var wrapper UserWrapper
	wrapper.Query = qryUser
	return repo.one(ctx, db, wrapper)
}

func (repo *User) FindMany(ctx context.Context, qryUser domain.User) domain.Users {
	return repo.findMany(ctx, repo.db, qryUser).Mapper()
}

func (repo *User) findMany(ctx context.Context, db *ent.Client, qryUser domain.User) ent.Users {
	var wrapper UserWrapper
	wrapper.Query = qryUser
	return repo.many(ctx, db, wrapper)
}

func (repo *User) One(ctx context.Context, wrapper UserWrapper) domain.User {
	return repo.one(ctx, repo.db, wrapper).Mapper()
}

func (repo *User) one(ctx context.Context, db *ent.Client, wrapper UserWrapper) *ent.User {
	builder := repo.query(db, wrapper)
	if wrapper.Option != nil {
		wrapper.Option(builder)
	}
	return builder.FirstX(ctx)
}

func (repo *User) Many(ctx context.Context, wrapper UserWrapper) domain.Users {
	var entUsers ent.Users = repo.many(ctx, repo.db, wrapper)
	return entUsers.Mapper()
}

func (repo *User) many(ctx context.Context, db *ent.Client, wrapper UserWrapper) []*ent.User {
	builder := repo.query(db, wrapper)
	if wrapper.Option != nil {
		wrapper.Option(builder)
	}
	return builder.AllX(ctx)
}

func (repo *User) query(db *ent.Client, wrapper UserWrapper) *ent.UserQuery {
	builder := db.User.Query()
	wrapper.Query.ID.IsPresent(func(v int64) { builder.Where(user.ID(v)) })
	wrapper.Query.Mobile.IsPresent(func(v string) { builder.Where(user.Mobile(v)) })
	wrapper.Query.Age.IsPresent(func(v int64) { builder.Where(user.Age(v)) })
	wrapper.Query.Age.StartIsPresent(func(v int64) { builder.Where(user.AgeGTE(v)) })
	wrapper.Query.Age.EndIsPresent(func(v int64) { builder.Where(user.AgeLTE(v)) })
	wrapper.Query.Level.IsPresent(func(v int64) { builder.Where(user.Level(v)) })
	wrapper.Query.Level.MultiIsPresent(func(v []int64) { builder.Where(user.LevelIn(v...)) })
	wrapper.Query.Nickname.IsPresent(func(v string) { builder.Where(user.Nickname(v)) })
	wrapper.Query.Nickname.FuzzyIsPresent(func(v string) { builder.Where(user.NicknameContains(v)) })
	wrapper.Query.CreateTime.StartIsPresent(func(v time.Time) { builder.Where(user.CreateTimeGTE(v)) })
	wrapper.Query.CreateTime.EndIsPresent(func(v time.Time) { builder.Where(user.CreateTimeLTE(v)) })
	return builder
}

func (repo *User) Save(ctx context.Context, cmdUser domain.User) domain.User {
	return repo.save(ctx, repo.db, cmdUser).Mapper()
}

func (repo *User) save(ctx context.Context, db *ent.Client, cmdUser domain.User) *ent.User {
	builder := db.User.Create()
	cmdUser.HashID.IsPresent(func(v string) { builder.SetHashID(v) })
	cmdUser.Mobile.IsPresent(func(v string) { builder.SetMobile(v) })
	cmdUser.Password.IsPresent(func(v string) { builder.SetPassword(v) })
	cmdUser.Age.IsPresent(func(v int64) { builder.SetAge(v) })
	cmdUser.Level.IsPresent(func(v int64) { builder.SetLevel(v) })
	cmdUser.Nickname.IsPresent(func(v string) { builder.SetNickname(v) })
	cmdUser.Avatar.IsPresent(func(v string) { builder.SetAvatar(v) })
	cmdUser.Bio.IsPresent(func(v string) { builder.SetBio(v) })
	cmdUser.Amount.IsPresent(func(v int64) { builder.SetAmount(v) })
	return builder.SaveX(ctx)
}

func (repo *User) Modify(ctx context.Context, cmdUser domain.User) int {
	return repo.modify(ctx, repo.db, cmdUser)
}

func (repo *User) modify(ctx context.Context, db *ent.Client, cmdUser domain.User) int {
	builder := db.User.Update()
	cmdUser.HashID.IsPresent(func(v string) { builder.SetHashID(v) })
	cmdUser.Mobile.IsPresent(func(v string) { builder.SetMobile(v) })
	cmdUser.Password.IsPresent(func(v string) { builder.SetPassword(v) })
	cmdUser.Age.IsPresent(func(v int64) { builder.SetAge(v) })
	cmdUser.Level.IsPresent(func(v int64) { builder.SetLevel(v) })
	cmdUser.Nickname.IsPresent(func(v string) { builder.SetNickname(v) })
	cmdUser.Avatar.IsPresent(func(v string) { builder.SetAvatar(v) })
	cmdUser.Bio.IsPresent(func(v string) { builder.SetBio(v) })
	cmdUser.Amount.IsPresent(func(v int64) { builder.SetAmount(v) })
	return builder.Where(user.ID(cmdUser.ID.Value)).SaveX(ctx)
}

func (repo *User) Remove(ctx context.Context, cmdUser domain.User) int {
	return repo.remove(ctx, repo.db, cmdUser)
}

func (repo *User) remove(ctx context.Context, db *ent.Client, cmdUser domain.User) int {
	builder := db.User.Delete()
	cmdUser.ID.IsPresent(func(v int64) { builder.Where(user.ID(v)) })
	cmdUser.HashID.IsPresent(func(v string) { builder.Where(user.HashID(v)) })
	cmdUser.Mobile.IsPresent(func(v string) { builder.Where(user.Mobile(v)) })
	cmdUser.Password.IsPresent(func(v string) { builder.Where(user.Password(v)) })
	cmdUser.Age.IsPresent(func(v int64) { builder.Where(user.Age(v)) })
	cmdUser.Level.IsPresent(func(v int64) { builder.Where(user.Level(v)) })
	cmdUser.Nickname.IsPresent(func(v string) { builder.Where(user.Nickname(v)) })
	cmdUser.Avatar.IsPresent(func(v string) { builder.Where(user.Avatar(v)) })
	cmdUser.Bio.IsPresent(func(v string) { builder.Where(user.Bio(v)) })
	cmdUser.Amount.IsPresent(func(v int64) { builder.Where(user.Amount(v)) })
	return builder.ExecX(ctx)
}

func (repo *User) Count(ctx context.Context, option func(opt *ent.UserQuery)) int {
	return repo.count(ctx, repo.db, option)
}

func (repo *User) count(ctx context.Context, db *ent.Client, option func(opt *ent.UserQuery)) int {
	builder := db.User.Query()
	option(builder)
	return builder.CountX(ctx)
}

func (repo *User) Exist(ctx context.Context, option func(opt *ent.UserQuery)) bool {
	return repo.exist(ctx, repo.db, option)
}

func (repo *User) NotExist(ctx context.Context, option func(opt *ent.UserQuery)) bool {
	return !repo.exist(ctx, repo.db, option)
}

func (repo *User) exist(ctx context.Context, db *ent.Client, option func(opt *ent.UserQuery)) bool {
	builder := db.User.Query()
	option(builder)
	return builder.ExistX(ctx)
}

func (repo *User) Create(ctx context.Context, option func(opt *ent.UserCreate)) domain.User {
	return repo.create(ctx, repo.db, option).Mapper()
}

func (repo *User) create(ctx context.Context, db *ent.Client, option func(opt *ent.UserCreate)) *ent.User {
	builder := db.User.Create()
	option(builder)
	return builder.SaveX(ctx)
}

func (repo *User) Update(ctx context.Context, option func(opt *ent.UserUpdate)) int {
	return repo.update(ctx, repo.db, option)
}

func (repo *User) update(ctx context.Context, db *ent.Client, option func(opt *ent.UserUpdate)) int {
	builder := db.User.Update()
	option(builder)
	return builder.SaveX(ctx)
}

func (repo *User) Delete(ctx context.Context, option func(opt *ent.UserDelete)) int {
	return repo.delete(ctx, repo.db, option)
}

func (repo *User) delete(ctx context.Context, db *ent.Client, option func(opt *ent.UserDelete)) int {
	builder := db.User.Delete()
	option(builder)
	return builder.ExecX(ctx)
}
