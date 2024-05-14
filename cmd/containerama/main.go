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
	app.Static("/", "frontend/dist")

	// session middleware
	store, err := sqlitestore.NewSqliteStoreFromConnection(db.DB, "sessions", "/", 3600, []byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		log.Fatal(err)
	}
	app.Use(session.Middleware(store))

	// API
	api := app.Group("/api")
	api.POST("/login", apihandlers.Login)
	api.GET("/logout", apihandlers.Logout)

	apicontainers := api.Group("/containers", middleware.ValidateSession)
	apicontainers.GET("", apihandlers.GetContainers)
	apicontainers.PUT("/:id", apihandlers.StopContainer)

	app.Start(os.Getenv("BIND_ADDR"))
}
