package models

import (
	"database/sql"
	"time"
)

type User struct {
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
	Username     string         `db:"username"`
	FirstName    string         `db:"first_name"`
	LastName     sql.NullString `db:"last_name"`
	Email        sql.NullString `db:"email"`
	PasswordHash string         `db:"password_hash"`
	Role         string         `db:"role"`
	ID           int            `db:"id"`
}
