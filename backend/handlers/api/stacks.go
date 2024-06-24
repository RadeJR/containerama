package api

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

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

func CreateStackFromFileHandler(c echo.Context) error {
	claims := c.Get("user").(*types.ZitadelClaims)
	userID := claims.Subject

	var data types.StackData
	err := c.Bind(&data)
	if err!=nil{
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse data")
	}
	
	path := "data/stacks/"+data.Name+"/docker-compose.yml"
	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		slog.Error("Error creating directories")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	err = os.WriteFile(path, []byte(data.Content), 0644)
	if err != nil {
		slog.Error("Error creating file", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	err = services.CreateStackFromFile(data.Name, userID, path)
	if err != nil {
		slog.Error("Error creating stack", "error", err)
		return err
	}
	
	return c.JSON(http.StatusOK, "success")
}
