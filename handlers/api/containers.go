package api

import (
	"net/http"

	"github.com/RadeJR/containerama/services"
	"github.com/labstack/echo/v4"
)

func GetContainers(c echo.Context) error {
	cont, err := services.GetContainers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, cont)
}

func StopContainer(c echo.Context) error {
	err := services.StopContainer(c.Param("id"))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
