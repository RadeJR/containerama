package middleware

import (
	"github.com/RadeJR/itcontainers/handler"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func CreateLocals(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := handler.CustomContext{Context: c, Locals: make(map[string]interface{})}
		return next(cc)
	}
}

func ValidateSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return c.Redirect(302, "/login")
		}

		if sess.IsNew {
			return c.Redirect(302, "/login")
		}

		c.(handler.CustomContext).Locals["role"] = sess.Values["role"]

		return next(c)
	}
}

func OnlyAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.(handler.CustomContext).Locals["role"] == "admin" {
			return next(c)
		} else {
			return c.String(403, "Unauthorized")
		}
	}
}
