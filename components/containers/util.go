package containers

import "github.com/docker/go-connections/nat"

func displayEnv(env []string) string {
	var envStr string
	for _, e := range env {
		envStr += e + "\n"
	}
	return envStr
}

func displayVolumes(volumes []string) string {
	var volumesStr string
	for _, v := range volumes {
		volumesStr += v + "\n"
	}
	return volumesStr
}

func displayPorts(ports nat.PortMap) string {
	var portsStr string
	for k, v := range ports {
		for i := 0; i < len(v); i++ {
			portsStr += k.Port() + ":" + v[i].HostPort + "/" + k.Proto() + "\n"
		}
	}
	return portsStr
}
