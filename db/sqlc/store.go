package db

import (
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/multierr"
)

type Store interface {
	Querier
	CreatePostTx(ctx context.Context, arg CreatePostParams) (PostResult, error)
}

//SQLStoreprovides all functions to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) WithTransaction(
	ctx context.Context, db *sql.DB, transaction func(qa *Queries) (txErr error),
) (err error) {
	var tx *sql.Tx

	if tx, err = db.BeginTx(ctx, nil); err != nil {
		return fmt.Errorf("could not start transcation: %w", err)
	}

	defer func() {
		v := recover()

		switch {
		case v != nil:
			_ = tx.Rollback()

			panic(v)
		case err != nil:
			if rbErr := tx.Rollback(); rbErr != nil {
				err = multierr.Combine(
					err,
					fmt.Errorf("could not rollback transaction: %w", rbErr),
				)
			}
		default:
			if err = tx.Commit(); err != nil {
				err = fmt.Errorf("could not commit transaction: %w", err)
			}
		}
	}()

	err = transaction(&Queries{tx})

	return
}

type PostResult struct {
	Post Post `json:"post"`
}

func (store *SQLStore) CreatePostTx(ctx context.Context, arg CreatePostParams) (PostResult, error) {
	var result PostResult
	err := store.WithTransaction(ctx, store.db, func(q *Queries) (txErr error) {
		var err error

		result.Post, err = q.CreatePost(ctx, arg)
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}
