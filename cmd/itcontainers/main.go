package main

import (
	"log"
	"os"
	"time"

	"github.com/RadeJR/itcontainers/db"
	"github.com/RadeJR/itcontainers/handlers"
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

	db, err := db.InitializeDB()
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
	loginHandler := handlers.LoginHandler{
		DB: db,
	}
	app.GET("/login", loginHandler.ShowLoginPage)
	app.POST("/login", loginHandler.Login)
	app.GET("/logout", loginHandler.Logout)

	// USER
	userHandler := handlers.UserHandler{
		DB: db,
	}
	app.GET("/users", userHandler.ShowUsers, middleware.ValidateSession, middleware.OnlyAdmin)
	app.POST("/users", userHandler.CreateUser, middleware.ValidateSession, middleware.OnlyAdmin)
	app.GET("/users/create", userHandler.CreateUserForm, middleware.ValidateSession, middleware.OnlyAdmin)

	// PAGES
	pageHandler := handlers.PageHandler{}
	app.GET("/", pageHandler.ShowBase, middleware.ValidateSession)
	app.GET("/containers", pageHandler.Containers, middleware.ValidateSession)
	app.GET("/networks", pageHandler.Networks, middleware.ValidateSession)

	// CONTAINERS
	dockerHandler := handlers.DockerHandler{
		Cli: cli,
	}
	app.GET("/containers", dockerHandler.GetContainers, middleware.ValidateSession)
	app.GET("/containers/create", dockerHandler.CreateContainerPage, middleware.ValidateSession)
	app.POST("/containers/create", dockerHandler.CreateContainer, middleware.ValidateSession)
	app.GET("/containers/stop/:id", dockerHandler.StopContainer, middleware.ValidateSession)
	app.GET("/containers/start/:id", dockerHandler.StartContainer, middleware.ValidateSession)
	app.GET("/containers/restart/:id", dockerHandler.RestartContainer, middleware.ValidateSession)
	app.GET("/containers/remove/:id", dockerHandler.RemoveContainer, middleware.ValidateSession)

	app.Start(":3000")
}
