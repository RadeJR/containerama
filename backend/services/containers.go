package services

import (
	"context"
	"encoding/binary"
	"errors"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	localtypes "github.com/RadeJR/containerama/types"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func GetContainer(id string, userID string, roles []string) (types.ContainerJSON, error) {
	cont, err := cli.ContainerInspect(context.Background(), id)
	if err != nil && !client.IsErrNotFound(err) {
		return types.ContainerJSON{}, err
	}
	if client.IsErrNotFound(err) {
		return types.ContainerJSON{}, errors.New("Not found")
	}
	
	for _,v := range roles {
		if v == "admin" {
			return cont, nil
		}
	}

	if cont.Config.Labels["owner"] == userID {
		return cont, nil
	}

	stacks, err := GetStacks(userID, roles)
	if err != nil {
		return types.ContainerJSON{}, err
	}

	for _, v := range stacks {
		if cont.Config.Labels["com.docker.compose.project"] == v.Name {
			return cont, nil
		}
	}

	return types.ContainerJSON{}, errors.New("Forbidden")
}

func GetContainers(userID string, roles []string) ([]types.Container, error) {
	var cont []types.Container
	var err error
	for _, v := range roles {
		if v == "admin" {
			cont, err = cli.ContainerList(context.Background(), container.ListOptions{All: true})
			if err != nil {
				return nil, err
			}
			return cont, nil
		}
	}

	filterArgs := filters.NewArgs()
	filterArgs.Add("label", "owner="+userID)
	// get containers that user owns
	cont, err = cli.ContainerList(context.Background(), container.ListOptions{All: true, Filters: filterArgs})
	if err != nil {
		return nil, err
	}

	// get containers that are owned by roles
	for _, v := range roles {
		filterArgs := filters.NewArgs()
		filterArgs.Add("label", "roles="+v)
		contFromRoles, err := cli.ContainerList(context.Background(), container.ListOptions{All: true, Filters: filterArgs})
		if err != nil {
			return nil, err
		}
		cont = append(cont, contFromRoles...)
	}

	stacks, err := GetStacks(userID, roles)
	if err != nil {
		return nil, err
	}
	var stackFilter filters.Args
	for _, v := range stacks {
		stackFilter = filters.NewArgs()
		stackFilter.Add("label", "com.docker.compose.project="+v.Name)
		stackCont, err := cli.ContainerList(context.Background(), container.ListOptions{All: true, Filters: stackFilter})
		if err != nil {
			slog.Error("Couldn't read stack containers", "error", err)
		}
		cont = append(cont, stackCont...)
	}

	return cont, nil
}

func StopContainer(id string) error {
	cli.ContainerStop(context.Background(), id, container.StopOptions{})
	statusCh, errCh := cli.ContainerWait(context.Background(), id, container.WaitConditionNotRunning)
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
	return nil
}

func CreateContainer(data localtypes.ContainerData, userID string) (string, error) {
	ctx := context.Background()

	reader, err := cli.ImagePull(ctx, data.Image, image.PullOptions{})
	if err != nil {
		return "", err
	}
	io.Copy(os.Stdout, reader)

	var ports nat.PortMap
	if data.Ports != "" {
		_, ports, err = nat.ParsePortSpecs(strings.Split(data.Ports, "\n"))
		if err != nil {
			return "", err
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

	labels := parseLabelString(data.Labels)
	if userID != "" {
		labels["owner"] = userID
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:           data.Image,
		Env:             env,
		Cmd:             cmd,
		Entrypoint:      entrypoint,
		Labels:          labels,
		NetworkDisabled: data.NetworkDisabled,
	}, &container.HostConfig{
		Binds:        volumes,
		PortBindings: ports,
	}, nil, nil, data.Name)
	if err != nil {
		RemoveContainer(resp.ID, true)
		return "", err
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", err
	}

	return resp.ID, nil
}

// Funciton that starts the container
func StartContainer(id string) error {
	if err := cli.ContainerStart(context.Background(), id, container.StartOptions{}); err != nil {
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
	if err := cli.ContainerRemove(context.Background(), id, container.RemoveOptions{
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

func EditContainer(id string, data localtypes.ContainerData, userID string) (string, error) {
	if err := StopContainer(id); err != nil {
		slog.Error("Error stopping container", "msg", err.Error())
		return "", err
	}
	if err := RemoveContainer(id, true); err != nil {
		slog.Error("Error removing container", "msg", err.Error())
		return "", err
	}
	id, err := CreateContainer(data, userID)
	if err != nil {
		slog.Error("Error creating container", "msg", err.Error())
		return "", err
	}
	return id, nil
}

func ContainerLogs(ctx context.Context, id string, logCh chan string) {
	defer close(logCh)

	reader, err := cli.ContainerLogs(ctx, id, container.LogsOptions{ShowStdout: true, ShowStderr: true, Follow: true, Tail: "all"})
	if err != nil {
		return
	}
	defer reader.Close()

	header := make([]byte, 8)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Read the 8-byte header
			_, err := reader.Read(header)
			if err != nil {
				return
			}

			// Parse the header
			streamType := header[0]
			var prefix string
			switch streamType {
			case 1:
				prefix = "OUT"
			default:
				prefix = "ERR"
			}
			payloadLength := binary.BigEndian.Uint32(header[4:8])

			// Read the payload
			payload := make([]byte, payloadLength)
			_, err = io.ReadFull(reader, payload)
			if err != nil {
				return
			}

			logCh <- prefix + string(payload)
		}
	}
}
