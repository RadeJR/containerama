package api

import (
	"net/http"

	"github.com/RadeJR/containerama/services"
	"github.com/labstack/echo/v4"
)

func ShowNetworks(c echo.Context) error {
	nr, err := services.GetNetworks()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nr)
}
