package main

import (
	"log"

	"github.com/RadeJR/itcontainers/db"
	"github.com/RadeJR/itcontainers/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env")
	}

	db, err := db.InitializeDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.User{})
}
