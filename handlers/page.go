package handlers

import (
	"github.com/RadeJR/itcontainers/components"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type PageHandler struct{}

func (h PageHandler) ShowBase(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(500, "Server error")
	}
	return render(c, components.Base(sess.Values["role"].(string)))
}

func (h PageHandler) Networks(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(500, "Server error")
	}
	return render(c, components.Networks(sess.Values["role"].(string)))
}
