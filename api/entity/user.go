package entity

import "time"

type User struct {
	ID        *int64    `db:"id"`
	Email     string    `db:"email"`
	UserName  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
