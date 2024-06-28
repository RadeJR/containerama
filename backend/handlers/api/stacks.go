package api

import (
	"log/slog"
	"net/http"

	"github.com/RadeJR/containerama/services"
	"github.com/RadeJR/containerama/types"
	"github.com/labstack/echo/v4"
)

func GetStacks(c echo.Context) error {
	claims := c.Get("user").(*types.ZitadelClaims)
	userID := claims.Subject
	isAdmin := false
	if claims.Roles["admin"] != nil {
		isAdmin = true
	}

	stacks, err := services.GetStacks(userID, isAdmin)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, stacks)
}

func CreateStack(c echo.Context) error {
	claims := c.Get("user").(*types.ZitadelClaims)
	userID := claims.Subject

	var data types.StackData
	err := c.Bind(&data)
	if err!=nil{
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
	claims := c.Get("user").(*types.ZitadelClaims)
	userID := claims.Subject

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
