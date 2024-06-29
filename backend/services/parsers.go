package services

import (
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
)

func parseLabelString(labelString string) map[string]string {
	result := make(map[string]string)

	lines := strings.Split(labelString, "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			result[key] = value
		}
	}

	return result
}

func portMapToString(portMap nat.PortMap) string {
	var sb strings.Builder
	for port, bindings := range portMap {
		for _, binding := range bindings {
			sb.WriteString(fmt.Sprintf("%s:%s:%s\n", binding.HostIP, binding.HostPort, port))
		}
	}
	return sb.String()
}

func volumeDataToString(container types.ContainerJSON) string {
	var sb strings.Builder
	for _, mount := range container.Mounts {
		sb.WriteString(fmt.Sprintf("%s:%s",
			mount.Source, mount.Destination))
	}
	return sb.String()
}

func labelsToString(labels map[string]string) string {
	var sb strings.Builder
	for key, value := range labels {
		sb.WriteString(fmt.Sprintf("%s=%s\n", key, value))
	}
	return sb.String()
}
