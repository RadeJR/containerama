package db

import (
	"embed"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS
var DB *sqlx.DB

func InitializeDB() {
	if _, err := os.Stat("./data"); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("./data", os.ModeDir)
		} else {
			panic(err)
		}
	}
	var err error
	DB, err = sqlx.Open("sqlite3", "./data/db.sqlite3")
	if err != nil {
		panic(err)
	}

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}
	if err := goose.Up(DB.DB, "migrations"); err != nil {
		panic(err)
	}
}

func CloseDB() {
	DB.Close()
}
