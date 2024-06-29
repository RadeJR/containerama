package types

type ContainerData struct {
	Image           string `json:"image" validate:"required"`
	Name            string `json:"name"`
	Env             string `json:"env"`
	Cmd             string `json:"cmd"`
	Ports           string `json:"exposedPorts"`
	Volumes         string `json:"volumes"`
	Entrypoint      string `json:"entrypoint"`
	Labels          string `json:"labels"`
	NetworkDisabled bool   `json:"networkDisabled"`
}
