package api

import (
	"log/slog"
	"net/http"

	"github.com/RadeJR/containerama/services"
	"github.com/RadeJR/containerama/types"
	"github.com/labstack/echo/v4"
)

func GetStacks(c echo.Context) error {
	userID := c.Get("userID").(string)
	roles := c.Get("roles").([]string)

	stacks, err := services.GetStacks(userID, roles)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, stacks)
}

func CreateStack(c echo.Context) error {
	userID := c.Get("userID").(string)

	var data types.StackData
	err := c.Bind(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse data")
	}

	err = services.CreateStack(data, userID)
	if err != nil {
		slog.Error("Error creating stack", "error", err)
		return err
	}

	return c.JSON(http.StatusOK, "success")
}

func DeleteStack(c echo.Context) error {
	userID := c.Get("userID").(string)

	name := c.QueryParam("name")
	if name == "" {
		return echo.ErrUnprocessableEntity
	}

	err := services.DeleteStack(name, userID)
	if err != nil {
		slog.Error("Error deleting stack", "error", err)
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusNoContent)
}
