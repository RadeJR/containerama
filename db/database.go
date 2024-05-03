package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func InitializeDB() {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	//
	// var err error
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	var err error
	DB, err = sqlx.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	DB.Close()
}
