package services

import "github.com/docker/docker/client"

var cli *client.Client

func InitializeCient() error {
	var err error
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	return nil
}

func CloseClient() error {
	err := cli.Close()
	if err != nil {
		return err
	}
	return nil
}
