package handlers

import (
	"github.com/RadeJR/itcontainers/components"
	"github.com/RadeJR/itcontainers/services"
	"github.com/labstack/echo/v4"
)

type DockerHandler struct{}

func (h DockerHandler) GetContainers(c echo.Context) error {
	cont, err := services.GetContainers()
	if err != nil {
		return render(c, components.ErrorPopup(err))
	}
	if c.Request().Header.Get("HX-Request") != "true" {
		return render(c, components.ContainersPageFull(cont, c.(CustomContext).Locals["role"].(string)))
	} else {
		return render(c, components.ContainersPage(cont, c.(CustomContext).Locals["role"].(string)))
	}
}

func (h DockerHandler) StopContainer(c echo.Context) error {
	err := services.StopContainer(c.Param("id"))
	if err != nil {
		render(c, components.ErrorPopup(err))
	}
	return h.GetContainers(c)
}

func (h DockerHandler) CreateContainerPage(c echo.Context) error {
	return render(c, components.ShowCreateForm())
}

func (h DockerHandler) CreateContainer(c echo.Context) error {
	var data services.ContainerData
	c.Bind(&data)

	err := services.CreateContainer(data)
	if err != nil {
		render(c, components.ErrorPopup(err))
		return h.CreateContainerPage(c)
	}
	return h.GetContainers(c)
}

func (h DockerHandler) StartContainer(c echo.Context) error {
	err := services.StartContainer(c.Param("id"))
	if err != nil {
		render(c, components.ErrorPopup(err))
	}
	return h.GetContainers(c)
}

func (h DockerHandler) RestartContainer(c echo.Context) error {
	err := services.RestartContainer(c.Param("id"))
	if err != nil {
		render(c, components.ErrorPopup(err))
	}
	return h.GetContainers(c)
}

func (h DockerHandler) RemoveContainer(c echo.Context) error {
	err := services.RemoveContainer(c.Param("id"), false)
	if err != nil {
		render(c, components.ErrorPopup(err))
	}
	return h.GetContainers(c)
}

func (h DockerHandler) ShowContainer(c echo.Context) error {
	cont, err := services.GetContainer(c.Param("id"))
	if err != nil {
		render(c, components.ErrorPopup(err))
		return h.ShowContainer(c)
	}
	return render(c, components.ShowContainerPage(cont))
}
