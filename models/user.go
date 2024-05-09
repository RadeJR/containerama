package models

import (
	"database/sql"
)

type User struct {
	Base
	Username     string         `db:"username"`
	FirstName    string         `db:"first_name"`
	LastName     sql.NullString `db:"last_name"`
	Email        sql.NullString `db:"email"`
	PasswordHash string         `db:"password_hash"`
	Role         string         `db:"role"`
}
