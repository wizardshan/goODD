package repository

import (
	"chapter10/domain"
	"chapter10/repository/ent"
	"chapter10/repository/ent/user"
	"context"
)

type User struct {
	repo
}

func NewUser(db *ent.Client) *User {
	repo := new(User)
	repo.db = db
	return repo
}

func (repo *User) Register(ctx context.Context, cmdUser domain.User) (domain.User, error) {
	var domUser domain.User
	err := repo.withTx(ctx, repo.db, func(tx *ent.Tx) error {
		db := tx.Client()
		domUser = repo.save(ctx, db, cmdUser).Mapper()

		domUser.HashID.ID = domUser.ID.Value
		domUser.HashID.Encode()
		repo.update(ctx, db, func(opt *ent.UserUpdate) {
			opt.SetHashID(domUser.HashID.Value).Where(user.ID(domUser.ID.Value))
		})
		return nil
	})
	return domUser, err
}

func (repo *User) FindOne(ctx context.Context, qryUser domain.User) domain.User {
	return repo.findOne(ctx, repo.db, qryUser).Mapper()
}

func (repo *User) findOne(ctx context.Context, db *ent.Client, qryUser domain.User) *ent.User {
	return repo.query(db, qryUser).FirstX(ctx)
}

func (repo *User) FindMany(ctx context.Context, qryUser domain.User) domain.Users {
	return repo.findMany(ctx, repo.db, qryUser).Mapper()
}

func (repo *User) findMany(ctx context.Context, db *ent.Client, qryUser domain.User) ent.Users {
	return repo.query(db, qryUser).AllX(ctx)
}

func (repo *User) query(db *ent.Client, qryUser domain.User) *ent.UserQuery {
	builder := db.User.Query()
	qryUser.ID.IsPresent(func(v int64) { builder.Where(user.ID(v)) })
	qryUser.Mobile.IsPresent(func(v string) { builder.Where(user.Mobile(v)) })
	qryUser.Age.IsPresent(func(v int64) { builder.Where(user.Age(v)) })
	qryUser.Level.IsPresent(func(v int64) { builder.Where(user.Level(v)) })
	qryUser.Nickname.IsPresent(func(v string) { builder.Where(user.Nickname(v)) })
	return builder
}

func (repo *User) Save(ctx context.Context, cmdUser domain.User) domain.User {
	return repo.save(ctx, repo.db, cmdUser).Mapper()
}

func (repo *User) save(ctx context.Context, db *ent.Client, cmdUser domain.User) *ent.User {
	builder := db.User.Create()
	cmdUser.Mobile.IsPresent(func(v string) { builder.SetMobile(v) })
	cmdUser.Age.IsPresent(func(v int64) { builder.SetAge(v) })
	cmdUser.Level.IsPresent(func(v int64) { builder.SetLevel(v) })
	cmdUser.Nickname.IsPresent(func(v string) { builder.SetNickname(v) })
	return builder.SaveX(ctx)
}

func (repo *User) Modify(ctx context.Context, cmdUser domain.User) int {
	return repo.modify(ctx, repo.db, cmdUser)
}

func (repo *User) modify(ctx context.Context, db *ent.Client, cmdUser domain.User) int {
	builder := db.User.Update()
	cmdUser.Mobile.IsPresent(func(v string) { builder.SetMobile(v) })
	cmdUser.Age.IsPresent(func(v int64) { builder.SetAge(v) })
	cmdUser.Level.IsPresent(func(v int64) { builder.SetLevel(v) })
	cmdUser.Nickname.IsPresent(func(v string) { builder.SetNickname(v) })
	return builder.Where(user.ID(cmdUser.ID.Value)).SaveX(ctx)
}

func (repo *User) Remove(ctx context.Context, cmdUser domain.User) int {
	return repo.remove(ctx, repo.db, cmdUser)
}

func (repo *User) remove(ctx context.Context, db *ent.Client, cmdUser domain.User) int {
	builder := db.User.Delete()
	cmdUser.ID.IsPresent(func(v int64) { builder.Where(user.ID(v)) })
	cmdUser.Mobile.IsPresent(func(v string) { builder.Where(user.Mobile(v)) })
	cmdUser.Age.IsPresent(func(v int64) { builder.Where(user.Age(v)) })
	cmdUser.Level.IsPresent(func(v int64) { builder.Where(user.Level(v)) })
	cmdUser.Nickname.IsPresent(func(v string) { builder.Where(user.Nickname(v)) })
	return builder.ExecX(ctx)
}

func (repo *User) Fetch(ctx context.Context, id int64) domain.User {
	return repo.FetchOne(ctx, func(opt *ent.UserQuery) {
		opt.Where(user.ID(id))
	})
}

func (repo *User) FetchOne(ctx context.Context, optionFunc func(opt *ent.UserQuery)) domain.User {
	return repo.fetchOne(ctx, repo.db, optionFunc).Mapper()
}

func (repo *User) fetchOne(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.UserQuery)) *ent.User {
	builder := db.User.Query()
	optionFunc(builder)
	return builder.FirstX(ctx)
}

func (repo *User) FetchMany(ctx context.Context, optionFunc func(opt *ent.UserQuery)) domain.Users {
	return repo.fetchMany(ctx, repo.db, optionFunc).Mapper()
}

func (repo *User) fetchMany(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.UserQuery)) ent.Users {
	builder := db.User.Query()
	optionFunc(builder)
	return builder.AllX(ctx)
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
