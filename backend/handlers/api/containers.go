package api

import (
	"context"
	"log/slog"
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
	return c.NoContent(http.StatusNoContent)
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

func ContainerLogs(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logCh := make(chan string, 100)
	slog.Info("SSE client connected", "ip", c.RealIP())

	w := c.Response()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	go services.ContainerLogs(ctx, c.Param("id"), logCh)

	for {
		select {
		case <-c.Request().Context().Done():
			slog.Info("SSE client disconnected", "ip", c.RealIP())
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
