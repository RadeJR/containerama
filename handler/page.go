package handler

import (
	"github.com/RadeJR/itcontainers/view/home"
	"github.com/RadeJR/itcontainers/view/layout"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type PageHandler struct{}

func (h PageHandler) ShowBase(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(500, "Server error")
	}
	return render(c, layout.Base(sess.Values["role"].(string)))
}

func (h PageHandler) ShowHome(c echo.Context) error {
	return render(c, home.Home())
}
