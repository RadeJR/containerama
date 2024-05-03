package main

import (
	"log"

	"github.com/RadeJR/containerama/db"
	"github.com/RadeJR/containerama/models"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env")
	}
	db.InitializeDB()
}

func main() {
	db.DB.AutoMigrate(&models.User{})
}
