package services

import (
	"context"

	"github.com/docker/docker/api/types"
)

func GetNetworks() ([]types.NetworkResource, error) {
  networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
  if err != nil {
    return nil, err
  }
  return networks, nil
}
