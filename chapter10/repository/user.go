package repository

import (
	"context"
	"goODD/chapter10/domain"
	"goODD/chapter10/repository/ent"
	"goODD/chapter10/repository/ent/user"
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
