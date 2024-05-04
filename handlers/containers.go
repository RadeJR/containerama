package handlers

import (
	"strconv"

	"github.com/RadeJR/containerama/components"
	"github.com/RadeJR/containerama/components/containers"
	"github.com/RadeJR/containerama/services"
	"github.com/labstack/echo/v4"
)

type DockerHandler struct{}

func (h DockerHandler) GetContainers(c echo.Context) error {
	pageString := c.QueryParam("page")
	var pageNum int
	if pageString != "" {
		var err error
		pageNum, err = strconv.Atoi(pageString)
		if err != nil {
			c.Response().Header().Set("HX-Retarget", "#popup")
			return Render(c, 500, components.ErrorPopup(err, false))
		}
	} else {
		pageNum = 1
	}
	sizeOfPageString := c.QueryParam("size")
	var sizeOfPageNum int
	if sizeOfPageString != "" {
		var err error
		sizeOfPageNum, err = strconv.Atoi(sizeOfPageString)
		if err != nil {
			c.Response().Header().Set("HX-Retarget", "#popup")
			return Render(c, 500, components.ErrorPopup(err, false))
		}
	} else {
		sizeOfPageNum = 10
	}
	cont, err := services.GetContainers()
	paginatedCont := PaginateContainers(cont, pageNum, sizeOfPageNum)
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return Render(c, 500, components.ErrorPopup(err, false))
	}
	if c.Request().Header.Get("HX-Request") != "true" {
		return Render(c, 200, containers.PageFull(paginatedCont, pageNum, sizeOfPageNum, len(cont), c.(CustomContext).Locals["role"].(string)))
	} else {
		role := c.(CustomContext).Locals["role"].(string)
		return Render(c, 200, containers.Page(paginatedCont, pageNum, sizeOfPageNum, len(cont), role))
	}
}

func (h DockerHandler) StopContainer(c echo.Context) error {
	err := services.StopContainer(c.Param("id"))
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return Render(c, 500, components.ErrorPopup(err, false))
	}
	return h.GetContainers(c)
}

func (h DockerHandler) CreateContainerPage(c echo.Context) error {
	return Render(c, 200, containers.Create())
}

func (h DockerHandler) CreateContainer(c echo.Context) error {
	var data services.ContainerData
	c.Bind(&data)

	err := services.CreateContainer(data)
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return Render(c, 500, components.ErrorPopup(err, false))
		// return h.CreateContainerPage(c)
	}
	return h.GetContainers(c)
}

func (h DockerHandler) StartContainer(c echo.Context) error {
	err := services.StartContainer(c.Param("id"))
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return Render(c, 200, components.ErrorPopup(err, false))
	}
	return h.GetContainers(c)
}

func (h DockerHandler) RestartContainer(c echo.Context) error {
	err := services.RestartContainer(c.Param("id"))
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return Render(c, 500, components.ErrorPopup(err, false))
	}
	return h.GetContainers(c)
}

func (h DockerHandler) RemoveContainer(c echo.Context) error {
	err := services.RemoveContainer(c.Param("id"), false)
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return Render(c, 500, components.ErrorPopup(err, false))
	}
	return h.GetContainers(c)
}

func (h DockerHandler) ShowContainer(c echo.Context) error {
	cont, err := services.GetContainer(c.Param("id"))
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return Render(c, 500, components.ErrorPopup(err, false))
	}
	return Render(c, 200, containers.One(cont))
}

func (h DockerHandler) EditContainerPage(c echo.Context) error {
	cont, err := services.GetContainer(c.Param("id"))
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return Render(c, 500, components.ErrorPopup(err, false))
	}
	return Render(c, 200, containers.Edit(cont))
}

func (h DockerHandler) EditContainer(c echo.Context) error {
	var data services.ContainerData
	c.Bind(&data)
	err := services.EditContainer(c.Param("id"), data)
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return Render(c, 500, components.ErrorPopup(err, false))
	}
	return h.GetContainers(c)
}
