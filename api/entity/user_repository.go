package entity

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	FindById(ctx context.Context, id int64) (*User, error)
	Create(ctx context.Context, tx *sqlx.Tx, user *User) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
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

func (r *DefaultUserRepository) Create(ctx context.Context, tx *sqlx.Tx, user *User) error {
	query := "INSERT INTO users (email, username, password) VALUES (:email, :username, :password)"

	// Execute the query and retrieve the result
	result, err := tx.NamedExecContext(ctx, query, user)
	if err != nil {
		return err
	}

	// Get the generated ID and bind it to the user struct
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = &id
	return nil
}

func (r *DefaultUserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int
	err := r.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM users WHERE email = ?", email)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
