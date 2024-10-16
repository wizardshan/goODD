package repository

import (
	"context"
	"goODD/chapter9/repository/ent"

	"github.com/pkg/errors"
)

type repo struct {
	db *ent.Client
}

func (repo *repo) withTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", errRollback)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}
