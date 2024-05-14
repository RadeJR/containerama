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

func CreateContainer(c echo.Context) error {
	var data services.ContainerData
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	err = services.Validate.Struct(data)
	if err != nil {
		return err
	}

	var id string
	id, err = services.CreateContainer(data)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, id)
}

func StartContainer(c echo.Context) error {
	err := services.StartContainer(c.Param("id"))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func RestartContainer(c echo.Context) error {
	err := services.RestartContainer(c.Param("id"))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func RemoveContainer(c echo.Context) error {
	err := services.RemoveContainer(c.Param("id"), false)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func ShowContainer(c echo.Context) error {
	cont, err := services.GetContainer(c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, cont)
}

func EditContainer(c echo.Context) error {
	var data services.ContainerData
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	err = services.Validate.Struct(data)
	if err != nil {
		return err
	}

	err = services.EditContainer(c.Param("id"), data)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, c.Param("id"))
}
