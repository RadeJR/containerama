package services

import (
	"context"
	"io"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

var Cli *client.Client

type ContainerData struct {
	Image           string `form:"image"`
	Name            string `form:"name"`
	Env             string `form:"env"`
	Cmd             string `form:"cmd"`
	Ports           string `form:"exposedPorts"`
	Volumes         string `form:"volumes"`
	Entrypoint      string `form:"entrypoint"`
	Labels          string `form:"labels"`
	NetworkDisabled bool   `form:"networkDisabled"`
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

	_, ports, err := nat.ParsePortSpecs(strings.Split(data.Ports, "\n"))
	if err != nil {
		return err
	}

	entrypoint := strings.Split(data.Entrypoint, " ")
	if entrypoint[0] == "" {
		entrypoint = nil
	}

	cmd := strings.Split(data.Cmd, " ")
	if cmd[0] == "" {
		cmd = nil
	}

	resp, err := Cli.ContainerCreate(ctx, &container.Config{
		Image:           data.Image,
		Env:             strings.Split(data.Env, "\n"),
		Cmd:             cmd,
		Entrypoint:      entrypoint,
		Labels:          parseLabelString(data.Labels),
		NetworkDisabled: data.NetworkDisabled,
	}, &container.HostConfig{
		Binds:        strings.Split(data.Volumes, "\n"),
		PortBindings: ports,
	}, nil, nil, data.Name)
	if err != nil {
		RemoveContainer(resp.ID, true)
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
func RemoveContainer(id string, force bool) error {
	statusCh, errCh := Cli.ContainerWait(context.Background(), id, container.WaitConditionNotRunning)
	if err := Cli.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{
		Force: force,
	}); err != nil {
		return err
	}

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
