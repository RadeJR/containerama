package handlers

import (
	"log/slog"

	"github.com/RadeJR/containerama/components"
	"github.com/RadeJR/containerama/db"
	"github.com/RadeJR/containerama/models"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct{}

func (h LoginHandler) ShowLoginPage(c echo.Context) error {
	return Render(c, 200, components.Login())
}

func (h LoginHandler) Login(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(500, err.Error())
	}

	if !sess.IsNew {
		return c.Redirect(302, "/")
	}

	type formData struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}
	var data formData
	c.Bind(&data)

	user := models.User{}

	err = db.DB.Get(&user, "SELECT * FROM users WHERE username = ?", data.Username)
	if err != nil {
		slog.Error("Failed getting user from db", "username", data.Username, "error", err)
		return c.String(504, "Wrong username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(data.Password)); err != nil {
		return c.String(504, "Wrong username or password")
	}

	sess.Values["name"] = user.Username
	sess.Values["role"] = user.Role

	sess.Save(c.Request(), c.Response())

	return c.Redirect(302, "/")
}

func (h LoginHandler) Logout(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(500, "server errror")
	}
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())
	return c.Redirect(302, "/login")
}
