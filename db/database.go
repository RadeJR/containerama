package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func InitializeDB() {
	var err error
	DB, err = sqlx.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	DB.Close()
}
