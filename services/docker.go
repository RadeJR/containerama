package services

import "github.com/docker/docker/client"

type imageError struct {
	error
	Message string
}

func (e imageError) Error() string {
	return e.Message
}

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
