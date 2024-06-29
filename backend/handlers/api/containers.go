package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/RadeJR/containerama/services"
	"github.com/RadeJR/containerama/types"
	dockertypes "github.com/docker/docker/api/types"
	"github.com/labstack/echo/v4"
)

func GetContainers(c echo.Context) error {
	userID := c.Get("userID").(string)
	roles := c.Get("roles").([]string)

	var cont []dockertypes.Container
	var err error
	cont, err = services.GetContainers(userID, roles)
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
	return c.NoContent(http.StatusNoContent)
}

func CreateContainer(c echo.Context) error {
	userID := c.Get("userID").(string)

	var data types.ContainerData
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	err = services.Validate.Struct(data)
	if err != nil {
		return err
	}

	var id string
	id, err = services.CreateContainer(data, userID)
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

func GetContainer(c echo.Context) error {
	userID := c.Get("userID").(string)
	roles := c.Get("roles").([]string)

	cont, err := services.GetContainer(c.Param("id"), userID, roles)
	if err != nil {
		slog.Error("Error getting container", "error", err)
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, cont)
}

func EditContainer(c echo.Context) error {
	userID := c.Get("userID").(string)
	
	var data types.ContainerData
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	err = services.Validate.Struct(data)
	if err != nil {
		return err
	}

	id, err := services.EditContainer(c.Param("id"), data, userID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, id)
}

func ContainerLogs(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logCh := make(chan string, 100)

	w := c.Response()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	go services.ContainerLogs(ctx, c.Param("id"), logCh)

	for {
		select {
		case <-c.Request().Context().Done():
			cancel()
			return nil
		case payload := <-logCh:
			event := services.Event{
				Data: []byte(payload),
			}
			if err := event.MarshalTo(w); err != nil {
				return err
			}
			w.Flush()
		}
	}
}
