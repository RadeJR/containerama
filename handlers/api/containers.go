package api

import (
	"net/http"

	"github.com/RadeJR/containerama/handlers"
	"github.com/RadeJR/containerama/services"
	"github.com/labstack/echo/v4"
)

func GetContainers(c echo.Context) error {
	page, size, err := handlers.GetPaginationInfo(c)
	if err != nil {
		return err
	}
	cont, err := services.GetContainers()
	if err != nil {
		return err
	}
	paginatedCont := services.PaginateContainers(cont, page, size)

	return c.JSON(http.StatusOK, paginatedCont)
}
