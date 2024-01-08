package main

import (
	"log"
	"os"
	"time"

	"github.com/RadeJR/itcontainers/database"
	"github.com/RadeJR/itcontainers/handler"
	"github.com/RadeJR/itcontainers/middleware"
	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/wader/gormstore/v2"
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

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	defer cli.Close()

	app := echo.New()
	// STATIC
	app.Static("/", "assets")

	app.Use(middleware.CreateLocals)

	// session middleware
	store := gormstore.New(db, []byte(os.Getenv("SESSION_SECRET")))
	quit := make(chan struct{})
	go store.PeriodicCleanup(1*time.Hour, quit)
	app.Use(session.Middleware(store))

	// LOGIN
	loginHandler := handler.LoginHandler{
		DB: db,
	}
	app.GET("/login", loginHandler.ShowLoginPage)
	app.POST("/login", loginHandler.Login)
	app.GET("/logout", loginHandler.Logout)

	// USER
	userHandler := handler.UserHandler{
		DB: db,
	}
	app.GET("/users", userHandler.ShowUsers, middleware.ValidateSession, middleware.OnlyAdmin)
	app.POST("/users", userHandler.CreateUser, middleware.ValidateSession, middleware.OnlyAdmin)
	app.GET("/users/create", userHandler.CreateUserForm, middleware.ValidateSession, middleware.OnlyAdmin)

	// PAGES
	pageHandler := handler.PageHandler{}
	app.GET("/", pageHandler.ShowBase, middleware.ValidateSession)
	app.GET("/containers", pageHandler.Containers, middleware.ValidateSession)
	app.GET("/networks", pageHandler.Networks, middleware.ValidateSession)

	// CONTAINERS
	dockerHandler := handler.DockerHandler{
		Cli: cli,
	}
	app.GET("/containers", dockerHandler.GetContainers, middleware.ValidateSession)

	app.Start(":3000")
}
