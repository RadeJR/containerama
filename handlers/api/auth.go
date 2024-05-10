package api

import (
	"net/http"

	"github.com/RadeJR/containerama/services"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	if !sess.IsNew {
		return c.JSON(http.StatusOK, "Already logged in")
	}

	type formData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var data formData
	err = c.Bind(&data)
	if err != nil {
		return err
	}

	user, err := services.GetUserByUsername(data.Username)
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

	return c.NoContent(http.StatusNoContent)
}

func Logout(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Options.MaxAge = -1
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
