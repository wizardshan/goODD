package repository

import (
	"context"
	"github.com/samber/lo"
	"goODD/chapter9/domain"
	"goODD/chapter9/repository/ent"
	"goODD/chapter9/repository/ent/user"
	"time"
)

type UserWrapper struct {
	Query  domain.User
	Option func(opt *ent.UserQuery)
	Fields domain.Fields
}

type User struct {
	repo
}

func NewUser(db *ent.Client) *User {
	repo := new(User)
	repo.db = db
	return repo
}

func (repo *User) Fetch(ctx context.Context, id int64) domain.User {
	return repo.fetch(ctx, repo.db, id).Mapper(nil)
}

func (repo *User) fetch(ctx context.Context, db *ent.Client, id int64) *ent.User {
	return repo.fetchOne(ctx, db, func(opt *ent.UserQuery) {
		opt.Where(user.ID(id))
	})
}

func (repo *User) FetchIn(ctx context.Context, ids ...int64) domain.User {
	return repo.fetchIn(ctx, repo.db, ids...).Mapper(nil)
}

func (repo *User) fetchIn(ctx context.Context, db *ent.Client, ids ...int64) *ent.User {
	return repo.fetchOne(ctx, db, func(opt *ent.UserQuery) {
		opt.Where(user.IDIn(ids...))
	})
}

func (repo *User) FetchOne(ctx context.Context, option func(opt *ent.UserQuery)) domain.User {
	return repo.fetchOne(ctx, repo.db, option).Mapper(nil)
}

func (repo *User) fetchOne(ctx context.Context, db *ent.Client, option func(opt *ent.UserQuery)) *ent.User {
	var wrapper UserWrapper
	wrapper.Option = option
	return repo.one(ctx, db, wrapper)
}

func (repo *User) FetchMany(ctx context.Context, option func(opt *ent.UserQuery)) domain.Users {
	return repo.fetchMany(ctx, repo.db, option).Mapper(nil)
}

func (repo *User) fetchMany(ctx context.Context, db *ent.Client, option func(opt *ent.UserQuery)) ent.Users {
	var wrapper UserWrapper
	wrapper.Option = option
	return repo.many(ctx, db, wrapper)
}

func (repo *User) FindOne(ctx context.Context, qryUser domain.User) domain.User {
	return repo.findOne(ctx, repo.db, qryUser).Mapper(nil)
}

func (repo *User) findOne(ctx context.Context, db *ent.Client, qryUser domain.User) *ent.User {
	var wrapper UserWrapper
	wrapper.Query = qryUser
	return repo.one(ctx, db, wrapper)
}

func (repo *User) FindMany(ctx context.Context, qryUser domain.User) domain.Users {
	return repo.findMany(ctx, repo.db, qryUser).Mapper(nil)
}

func (repo *User) findMany(ctx context.Context, db *ent.Client, qryUser domain.User) ent.Users {
	var wrapper UserWrapper
	wrapper.Query = qryUser
	return repo.many(ctx, db, wrapper)
}

func (repo *User) One(ctx context.Context, wrapper UserWrapper) domain.User {
	return repo.one(ctx, repo.db, wrapper).Mapper(wrapper.Fields)
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
	return entUsers.Mapper(wrapper.Fields)
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
	if names, ok := wrapper.Fields.SnakeCaseNames(user.Label); ok {
		builder.Select(lo.Intersect(user.Columns, names)...)
	}
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
