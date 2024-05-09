package handlers

import (
	"net/http"

	"github.com/RadeJR/containerama/components"
	"github.com/RadeJR/containerama/db"
	"github.com/RadeJR/containerama/models"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct{}

func (h LoginHandler) ShowLoginPage(c echo.Context) error {
	if c.Request().Header.Get("HX-Request") != "true" {
		return Render(c, 200, components.LoginPage())
	} else {
		return Render(c, 200, components.Login())
	}
}

func (h LoginHandler) Login(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	if !sess.IsNew {
		return c.Redirect(302, "/")
	}

	type formData struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}
	var data formData
	err = c.Bind(&data)
	if err != nil {
		return err
	}

	user := models.User{}

	err = db.DB.Get(&user, "SELECT * FROM users WHERE username = ?", data.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Wrong username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(data.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Wrong username or password")
	}

	sess.Values["id"] = user.ID
	sess.Values["username"] = user.Username
	sess.Values["role"] = user.Role

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.Redirect(302, "/")
}

func (h LoginHandler) Logout(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Options.MaxAge = -1
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return c.Redirect(302, "/login")
}
