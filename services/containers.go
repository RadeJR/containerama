package services

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var Cli *client.Client

type ContainerData struct {
	Image string `form:"image"`
	Name  string `form:"name"`
}

func InitializeCient() error {
	var err error
	Cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	return nil
}

func GetContainers() ([]types.Container, error) {
	cont, err := Cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}
	return cont, nil
}

func StopContainer(id string) error {
	statusCh, errCh := Cli.ContainerWait(context.Background(), id, container.WaitConditionNotRunning)
	Cli.ContainerStop(context.Background(), id, container.StopOptions{})
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case status := <-statusCh:
		if status.Error != nil {
			return nil
		}
	}
	time.Sleep(time.Second)
	return nil
}

func CreateContainer(data ContainerData) error {
	ctx := context.Background()

	reader, err := Cli.ImagePull(ctx, data.Image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)

	resp, err := Cli.ContainerCreate(ctx, &container.Config{
		Image: data.Image,
	}, nil, nil, nil, data.Name)
	if err != nil {
		return err
	}

	if err := Cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	return nil
}

// Funciton that starts the container
func StartContainer(id string) error {
	if err := Cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{}); err != nil {
		return err
	}
	return nil
}

func RestartContainer(id string) error {
	if err := Cli.ContainerRestart(context.Background(), id, container.StopOptions{}); err != nil {
		return err
	}
	return nil
}

// function that deletes the container
func RemoveContainer(id string) error {
	if err := Cli.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{}); err != nil {
		return err
	}
	statusCh, errCh := Cli.ContainerWait(context.Background(), id, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case status := <-statusCh:
		if status.StatusCode != 0 {
			return nil
		}
	}
	time.Sleep(time.Second)
	return nil
}

func GetContainer(id string) (types.ContainerJSON, error) {
	cont, err := Cli.ContainerInspect(context.Background(), id)
	if err != nil {
		return types.ContainerJSON{}, err
	}
	return cont, nil
}
