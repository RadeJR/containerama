package middleware

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func ValidateSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return c.Redirect(302, "/login")
		}

		if sess.IsNew {
			return c.Redirect(302, "/login")
		}

		return next(c)
	}
}

func OnlyAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return c.Redirect(302, "/login")
		}

		if sess.Values["role"] == "admin" {
			return next(c)
		} else {
			return c.String(403, "Unauthorized")
		}
	}
}
