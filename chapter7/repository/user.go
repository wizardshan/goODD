package repository

import (
	"context"
	"goODD/chapter7/domain"
	"goODD/chapter7/repository/ent"
	"goODD/chapter7/repository/ent/user"
)

type User struct {
	db *ent.Client
}

func NewUser(db *ent.Client) *User {
	repo := new(User)
	repo.db = db
	return repo
}

func (repo *User) FetchWithFields(ctx context.Context, id int64, fields domain.Fields) domain.User {
	query := repo.db.User.Query()
	names, ok := fields.SnakeCaseNames(user.Label)
	if ok {
		query.Select(names...)
	}
	return query.Where(user.ID(id)).FirstX(ctx).MapperWithFields(fields)
}

func (repo *User) Fetch(ctx context.Context, qryUser domain.User) domain.User {
	return repo.db.User.Query().Where(user.ID(qryUser.ID.Value)).FirstX(ctx).Mapper()
}

func (repo *User) FetchOne(ctx context.Context, qryUser domain.User) domain.User {
	return repo.query(qryUser).FirstX(ctx).Mapper()
}

func (repo *User) FetchMany(ctx context.Context, qryUser domain.User) domain.Users {
	var entUsers ent.Users = repo.query(qryUser).AllX(ctx)
	return entUsers.Mapper()
}

func (repo *User) query(qryUser domain.User) *ent.UserQuery {
	query := repo.db.User.Query()
	qryUser.ID.IsPresent(func(v int64) { query.Where(user.ID(v)) })
	qryUser.HashID.IsPresent(func(v string) { query.Where(user.HashID(v)) })
	qryUser.Mobile.IsPresent(func(v string) { query.Where(user.Mobile(v)) })
	qryUser.Password.IsPresent(func(v string) { query.Where(user.Password(v)) })
	qryUser.Age.IsPresent(func(v int64) { query.Where(user.Age(v)) })
	qryUser.Level.IsPresent(func(v int64, _ string) { query.Where(user.Level(v)) })
	qryUser.Nickname.IsPresent(func(v string) { query.Where(user.Nickname(v)) })
	qryUser.Avatar.IsPresent(func(v string) { query.Where(user.Avatar(v)) })
	qryUser.Bio.IsPresent(func(v string) { query.Where(user.Bio(v)) })
	qryUser.Amount.IsPresent(func(v int64) { query.Where(user.Amount(v)) })

	return query
}
