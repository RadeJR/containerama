package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/RadeJR/containerama/db"
	"github.com/RadeJR/containerama/handlers"
	apihandlers "github.com/RadeJR/containerama/handlers/api"
	"github.com/RadeJR/containerama/handlers/auth"
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
	auth.InitializeOauth()
	middleware.InitializeKeyFunc()
}

func main() {
	defer db.CloseDB()
	defer services.CloseClient()

	app := echo.New()
	app.HTTPErrorHandler = handlers.CustomHTTPErrorHandler
	// STATIC
	if os.Getenv("APP_ENV") == "local" {
		app.Static("/", "../frontend/dist")
	} else {
		app.Static("/", "public")
	}

	// session middleware
	store, err := sqlitestore.NewSqliteStoreFromConnection(db.DB, "sessions", "/", 3600, []byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		log.Fatal(err)
	}
	app.Use(session.Middleware(store))

	app.GET("/login", auth.LoginHandler)
	app.GET("/callback", auth.CallbackHandler)
	app.GET("/logout", auth.LogoutHandler)
	app.GET("/webhook/:id", apihandlers.StackWebhook)

	// API
	api := app.Group("/api", middleware.JWTMiddleware)
	api.GET("/userinfo", auth.LoginCheckHandler)

	apicontainers := api.Group("/containers", middleware.JWTMiddleware)
	apicontainers.GET("", apihandlers.GetContainers)
	apicontainers.GET("/:id", apihandlers.GetContainer)
	apicontainers.PUT("/:id/stop", apihandlers.StopContainer)
	apicontainers.PUT("/:id/start", apihandlers.StartContainer)
	apicontainers.DELETE("/:id", apihandlers.RemoveContainer)
	apicontainers.GET("/:id/logs", apihandlers.ContainerLogs)
	apicontainers.PATCH("/:id", apihandlers.EditContainer)
	apicontainers.POST("/create", apihandlers.CreateContainer)

	apistacks := api.Group("/stacks", middleware.JWTMiddleware)
	apistacks.GET("", apihandlers.GetStacks)
	apistacks.POST("", apihandlers.CreateStack)
	apistacks.DELETE("", apihandlers.DeleteStack)

	app.Start(os.Getenv("BIND_ADDR"))
}
