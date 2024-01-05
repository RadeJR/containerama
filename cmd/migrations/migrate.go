package main

import (
	"log"

	"github.com/RadeJR/itcontainers/database"
	"github.com/RadeJR/itcontainers/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env")
	}

	db, err := database.InitializeDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&model.User{})
}
