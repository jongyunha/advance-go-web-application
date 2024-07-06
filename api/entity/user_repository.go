package entity

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	FindById(ctx context.Context, id int64) (*User, error)
}

type DefaultUserRepository struct {
	db *sqlx.DB
}

func NewDefaultUserRepository(db *sqlx.DB) *DefaultUserRepository {
	return &DefaultUserRepository{db: db}
}

func (r *DefaultUserRepository) FindById(ctx context.Context, id int64) (*User, error) {
	var user User
	err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
