package middleware

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func ValidateSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil || sess.IsNew {
			if ok, _ := regexp.Match("/api/*", []byte(c.Path())); ok {
				sess.Options.MaxAge = -1
				if err := sess.Save(c.Request(), c.Response()); err != nil {
					return err
				}
				return c.JSON(http.StatusUnauthorized, "unauthorized")
			}
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
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
	}
}
