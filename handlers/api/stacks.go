package api

import (
	"net/http"

	"github.com/RadeJR/containerama/models"
	"github.com/RadeJR/containerama/services"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetStacks(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	role := sess.Values["role"].(string)
	id := sess.Values["id"].(int)

	stacks := []models.Stack{}
	stacks, err = services.GetStacks(id, role)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, stacks)
}
