package services

import (
	"context"
	"time"

	"github.com/RadeJR/containerama/components"
	"github.com/docker/docker/api/types"
)

func GetNetworks() ([]types.NetworkResource, error) {
	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return nil, err
	}
	return networks, nil
}

func NewNetworkRowData(nr types.NetworkResource) components.RowData {
	rowData := components.RowData{
		Fields: make([]string, 4),
	}

	rowData.Fields[0] = nr.ID[:10]
	rowData.Fields[1] = nr.Driver
	rowData.Fields[2] = nr.Name
	rowData.Fields[3] = nr.Created.Format(time.RFC3339)
	return rowData
}
