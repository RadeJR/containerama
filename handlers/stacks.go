package handlers

import (
	"github.com/RadeJR/containerama/services"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type StackHandler struct{}

func (sh StackHandler) GetStacks(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	role := sess.Values["role"].(string)
	id := sess.Values["id"].(int)

	page, size, err := GetPaginationInfo(c)
	if err != nil {
		return err
	}

	_, err = services.GetStacks(id, role, page, size)
	if err != nil {
		return err
	}
	return nil
}
