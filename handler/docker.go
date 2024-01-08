package handler

import (
	"context"

	"github.com/RadeJR/itcontainers/view/containers"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
)

type DockerHandler struct {
	Cli *client.Client
}

func (h DockerHandler) GetContainers(c echo.Context) error {
	conts, err := h.Cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return c.String(500, "Couldnt fetch containers")
	}

	return render(c, containers.ContainersPage(conts, c.(CustomContext).Locals["role"].(string)))
}
