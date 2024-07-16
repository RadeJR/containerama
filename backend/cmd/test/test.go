package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func loggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Request received")
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Use(loggingMiddleware)
	e.GET("/hello", helloHandler)
	e.Start(":8080")
}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

