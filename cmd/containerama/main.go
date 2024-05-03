package main

import (
	"log"
	"os"

	"github.com/RadeJR/containerama/db"
	"github.com/RadeJR/containerama/handlers"
	"github.com/RadeJR/containerama/middleware"
	"github.com/RadeJR/containerama/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/michaeljs1990/sqlitestore"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env")
	}

	db.InitializeDB()
	defer db.CloseDB()

	services.InitializeCient()
	defer services.CloseClient()
}

func main() {
	app := echo.New()
	// STATIC
	app.Static("/", "assets")

	app.Use(middleware.CreateLocals)

	// session middleware
	store, err := sqlitestore.NewSqliteStore("./db.sqlite3", "sessions", "/", 3600, []byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		log.Fatal(err)
	}
	app.Use(session.Middleware(store))

	// LOGIN
	loginHandler := handlers.LoginHandler{}
	app.GET("/login", loginHandler.ShowLoginPage)
	app.POST("/login", loginHandler.Login)
	app.GET("/logout", loginHandler.Logout)

	// USER
	userHandler := handlers.UserHandler{}
	users := app.Group("/users", middleware.ValidateSession, middleware.OnlyAdmin)
	users.GET("", userHandler.ShowUsers)
	users.POST("", userHandler.CreateUser)
	users.GET("/create", userHandler.CreateUserForm)

	// DOCKER
	dockerHandler := handlers.DockerHandler{}
	// Containers is default page
	app.GET("/", dockerHandler.GetContainers, middleware.ValidateSession)

	// CONTAINERS
	containers := app.Group("/containers", middleware.ValidateSession)
	containers.GET("", dockerHandler.GetContainers)
	containers.GET("/create", dockerHandler.CreateContainerPage)
	containers.POST("/create", dockerHandler.CreateContainer)
	containers.GET("/stop/:id", dockerHandler.StopContainer)
	containers.GET("/start/:id", dockerHandler.StartContainer)
	containers.GET("/restart/:id", dockerHandler.RestartContainer)
	containers.GET("/remove/:id", dockerHandler.RemoveContainer)
	containers.GET(":id", dockerHandler.ShowContainer)
	containers.GET("/edit/:id", dockerHandler.EditContainerPage)
	containers.POST("/edit/:id", dockerHandler.EditContainer)
	// NETWORKS
	// networks := app.Group("/networks", middleware.ValidateSession)
	// networks.GET("", dockerHandler.GetNetworks, middleware.ValidateSession)
	// networks.GET("/create", dockerHandler.CreateNetworkPage)
	// networks.POST("/create", dockerHandler.CreateNetwork)
	// networks.GET("/remove/:id", dockerHandler.RemoveNetwork)
	// networks.GET(":id", dockerHandler.ShowNetwork)

	app.Start(":3000")
}
