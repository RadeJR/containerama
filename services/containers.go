package services

import (
	"context"
	"io"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)


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

func GetContainers() ([]types.Container, error) {
	cont, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}
	return cont, nil
}

func StopContainer(id string) error {
	statusCh, errCh := cli.ContainerWait(context.Background(), id, container.WaitConditionNotRunning)
	cli.ContainerStop(context.Background(), id, container.StopOptions{})
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

	if data.Image == "" {
		return imageError{Message: "Image is required"}
	}

	reader, err := cli.ImagePull(ctx, data.Image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)

	var ports nat.PortMap
	if data.Ports != "" {
		_, ports, err = nat.ParsePortSpecs(strings.Split(data.Ports, "\n"))
		if err != nil {
			return err
		}
	} else {
		ports = nil
	}

	var entrypoint []string
	if data.Entrypoint != "" {
		entrypoint = strings.Split(data.Entrypoint, " ")
	} else {
		entrypoint = nil
	}

	var cmd []string
	if data.Cmd != "" {
		cmd = strings.Split(data.Cmd, " ")
	} else {
		cmd = nil
	}

	var env []string
	if data.Env != "" {
		env = strings.Split(data.Env, "\n")
	} else {
		env = nil
	}

	var volumes []string
	if data.Volumes != "" {
		volumes = strings.Split(data.Volumes, "\n")
	} else {
		volumes = nil
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:           data.Image,
		Env:             env,
		Cmd:             cmd,
		Entrypoint:      entrypoint,
		Labels:          parseLabelString(data.Labels),
		NetworkDisabled: data.NetworkDisabled,
	}, &container.HostConfig{
		Binds:        volumes,
		PortBindings: ports,
	}, nil, nil, data.Name)
	if err != nil {
		RemoveContainer(resp.ID, true)
		return err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	return nil
}

// Funciton that starts the container
func StartContainer(id string) error {
	if err := cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{}); err != nil {
		return err
	}
	return nil
}

func RestartContainer(id string) error {
	if err := cli.ContainerRestart(context.Background(), id, container.StopOptions{}); err != nil {
		return err
	}
	return nil
}

// function that deletes the container
func RemoveContainer(id string, force bool) error {
	statusCh, errCh := cli.ContainerWait(context.Background(), id, container.WaitConditionNotRunning)
	if err := cli.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{
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
	cont, err := cli.ContainerInspect(context.Background(), id)
	if err != nil {
		return types.ContainerJSON{}, err
	}
	return cont, nil
}

func EditContainer(id string, data ContainerData) error {
	if err := StopContainer(id); err != nil {
		return err
	}
	if err := RemoveContainer(id, true); err != nil {
		return err
	}
	if err := CreateContainer(data); err != nil {
		return err
	}
	return nil
}
