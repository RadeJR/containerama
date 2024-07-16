package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/RadeJR/containerama/services"
	"github.com/RadeJR/containerama/types"
	"github.com/go-playground/validator/v10"
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

	err = services.Validate.Struct(data)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		var msg []string
		for _, err := range err.(validator.ValidationErrors) {
			if err.Param() != "" {
				msg = append(msg, fmt.Sprintf("%v is required without %v.", err.Field(), err.Param()))
			} else {
				msg = append(msg, fmt.Sprintf("%v is required.", err.Field()))
			}
		}
		return echo.NewHTTPError(http.StatusUnprocessableEntity, msg)
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

func StackWebhook(c echo.Context) error {
	services.StackWebhook(c.Param("id"))
	return c.NoContent(http.StatusNoContent)
}
