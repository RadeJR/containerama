package models

import "time"

type User struct {
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	Username     string    `db:"username"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	Role         string    `db:"role"`
	ID           int       `db:"id"`
}
