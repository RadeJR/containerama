package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/RadeJR/containerama/db"
	"github.com/RadeJR/containerama/handlers"
	apihandlers "github.com/RadeJR/containerama/handlers/api"
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
		slog.Error("Failed to load env", "error", err)
	}
	db.InitializeDB()
	services.InitializeCient()
	services.InitializeValidator()
	services.EnsureAdminUserExists()
}

func main() {
	defer db.CloseDB()
	defer services.CloseClient()
	app := echo.New()
	app.HTTPErrorHandler = handlers.CustomHTTPErrorHandler
	// STATIC
	app.Static("/", "assets")

	// session middleware
	store, err := sqlitestore.NewSqliteStoreFromConnection(db.DB, "sessions", "/", 3600, []byte(os.Getenv("SESSION_SECRET")))
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
	networkHandler := handlers.NetworkHandler{}
	networks := app.Group("/networks", middleware.ValidateSession)
	networks.GET("", networkHandler.ShowNetworks, middleware.ValidateSession)
	// networks.GET("/create", dockerHandler.CreateNetworkPage)
	// networks.POST("/create", dockerHandler.CreateNetwork)
	// networks.GET("/remove/:id", dockerHandler.RemoveNetwork)
	// networks.GET(":id", dockerHandler.ShowNetwork)
	
	// API
	api := app.Group("/api")
	api.POST("/login", apihandlers.Login)
	api.GET("/logout", apihandlers.Logout)

	apicontainers := api.Group("/containers", middleware.ValidateSession)
	apicontainers.GET("", apihandlers.GetContainers)

	app.Start(os.Getenv("BIND_ADDR"))
}
