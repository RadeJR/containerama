package handler

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/RadeJR/itcontainers/view/containers"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
)

type DockerHandler struct {
	Cli *client.Client
}

func (h DockerHandler) GetContainers(c echo.Context) error {
	cont, err := h.Cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return c.String(500, "Couldnt fetch containers")
	}

	return render(c, containers.ContainersPage(cont, c.(CustomContext).Locals["role"].(string)))
}

func (h DockerHandler) StopContainer(c echo.Context) error {
	statusCh, errCh := h.Cli.ContainerWait(context.Background(), c.Param("id"), container.WaitConditionNotRunning)
	h.Cli.ContainerStop(context.Background(), c.Param("id"), container.StopOptions{})
	select {
	case err := <-errCh:
		if err != nil {
			return c.String(500, err.Error())
		}
	case status := <-statusCh:
		if status.StatusCode != 0 {
			return c.String(500, fmt.Sprintf("container exited with status %d", status.StatusCode))
		}
	}
	time.Sleep(time.Second)

	return c.Redirect(302, "/containers")
}

func (h DockerHandler) CreateContainerPage(c echo.Context) error {
	return render(c, containers.ShowCreateForm())
}

func (h DockerHandler) CreateContainer(c echo.Context) error {
	type ContainerData struct {
		Image string `form:"image"`
		Name  string `form:"name"`
	}
	var data ContainerData

	c.Bind(&data)

	ctx := context.Background()
	reader, err := h.Cli.ImagePull(ctx, data.Image, types.ImagePullOptions{})
	if err != nil {
		return c.String(500, err.Error())
	}
	io.Copy(os.Stdout, reader)

	resp, err := h.Cli.ContainerCreate(ctx, &container.Config{
		Image: data.Image,
	}, nil, nil, nil, data.Name)
	if err != nil {
		return c.String(500, err.Error())
	}

	if err := h.Cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return c.String(500, err.Error())
	}
	return c.Redirect(302, "/containers")
}
