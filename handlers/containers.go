package handlers

import (
	"github.com/RadeJR/itcontainers/components"
	"github.com/RadeJR/itcontainers/components/containers"
	"github.com/RadeJR/itcontainers/services"
	"github.com/labstack/echo/v4"
)

type DockerHandler struct{}

func (h DockerHandler) GetContainers(c echo.Context) error {
	cont, err := services.GetContainers()
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return render(c, components.ErrorPopup(err, false))
	}
	if c.Request().Header.Get("HX-Request") != "true" {
		return render(c, containers.PageFull(cont, c.(CustomContext).Locals["role"].(string)))
	} else {
    role := c.(CustomContext).Locals["role"].(string)
    render(c, components.Navbar(role, "Containers"))
		return render(c, containers.All(cont))
	}
}

func (h DockerHandler) StopContainer(c echo.Context) error {
	err := services.StopContainer(c.Param("id"))
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return render(c, components.ErrorPopup(err, false))
	}
	return h.GetContainers(c)
}

func (h DockerHandler) CreateContainerPage(c echo.Context) error {
	return render(c, containers.Create())
}

func (h DockerHandler) CreateContainer(c echo.Context) error {
	var data services.ContainerData
	c.Bind(&data)

	err := services.CreateContainer(data)
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return render(c, components.ErrorPopup(err, false))
		// return h.CreateContainerPage(c)
	}
	return h.GetContainers(c)
}

func (h DockerHandler) StartContainer(c echo.Context) error {
	err := services.StartContainer(c.Param("id"))
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return render(c, components.ErrorPopup(err, false))
	}
	return h.GetContainers(c)
}

func (h DockerHandler) RestartContainer(c echo.Context) error {
	err := services.RestartContainer(c.Param("id"))
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return render(c, components.ErrorPopup(err, false))
	}
	return h.GetContainers(c)
}

func (h DockerHandler) RemoveContainer(c echo.Context) error {
	err := services.RemoveContainer(c.Param("id"), false)
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return render(c, components.ErrorPopup(err, false))
	}
	return h.GetContainers(c)
}

func (h DockerHandler) ShowContainer(c echo.Context) error {
	cont, err := services.GetContainer(c.Param("id"))
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return render(c, components.ErrorPopup(err, false))
	}
	return render(c, containers.One(cont))
}

func (h DockerHandler) EditContainerPage(c echo.Context) error {
	cont, err := services.GetContainer(c.Param("id"))
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return render(c, components.ErrorPopup(err, false))
	}
	return render(c, containers.Edit(cont))
}

func (h DockerHandler) EditContainer(c echo.Context) error {
	var data services.ContainerData
	c.Bind(&data)
	err := services.EditContainer(c.Param("id"), data)
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return render(c, components.ErrorPopup(err, false))
	}
	return h.GetContainers(c)
}
