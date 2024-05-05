package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

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
			return RenderError(c, http.StatusBadRequest, err)
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
			return RenderError(c, http.StatusBadRequest, err)
		}
	} else {
		sizeOfPageNum = 10
	}
	cont, err := services.GetContainers()
	if err != nil {
		slog.Error("Error getting containers", "error", err.Error())
		return c.String(http.StatusInternalServerError, "Internal server error")
	}
	paginatedCont := PaginateContainers(cont, pageNum, sizeOfPageNum)
	role := c.(CustomContext).Locals["role"].(string)
	if c.Request().Header.Get("HX-Request") != "true" {
		return Render(c, 200, containers.PageFull(paginatedCont, pageNum, sizeOfPageNum, len(cont), role))
	} else {
		return Render(c, 200, containers.Page(paginatedCont, pageNum, sizeOfPageNum, len(cont), role))
	}
}

func (h DockerHandler) StopContainer(c echo.Context) error {
	err := services.StopContainer(c.Param("id"))
	if err != nil {
		return RenderDockerError(c, err)
	}
	return h.GetContainers(c)
}

func (h DockerHandler) CreateContainerPage(c echo.Context) error {
	return Render(c, 200, containers.Create())
}

func (h DockerHandler) CreateContainer(c echo.Context) error {
	var data services.ContainerData
	c.Bind(&data)

	var err error
	err = services.Validate.Struct(data)
	if err != nil {
		return RenderError(c, http.StatusBadRequest, err)
	}

	err = services.CreateContainer(data)
	if err != nil {
		return RenderDockerError(c, err)
	}
	return h.GetContainers(c)
}

func (h DockerHandler) StartContainer(c echo.Context) error {
	err := services.StartContainer(c.Param("id"))
	if err != nil {
		return RenderDockerError(c, err)
	}
	return h.GetContainers(c)
}

func (h DockerHandler) RestartContainer(c echo.Context) error {
	err := services.RestartContainer(c.Param("id"))
	if err != nil {
		return RenderDockerError(c, err)
	}
	return h.GetContainers(c)
}

func (h DockerHandler) RemoveContainer(c echo.Context) error {
	err := services.RemoveContainer(c.Param("id"), false)
	if err != nil {
		return RenderDockerError(c, err)
	}
	return h.GetContainers(c)
}

func (h DockerHandler) ShowContainer(c echo.Context) error {
	cont, err := services.GetContainer(c.Param("id"))
	if err != nil {
		return RenderDockerError(c, err)
	}
	return Render(c, 200, containers.One(cont))
}

func (h DockerHandler) EditContainerPage(c echo.Context) error {
	cont, err := services.GetContainer(c.Param("id"))
	if err != nil {
		return RenderDockerError(c, err)
	}
	return Render(c, 200, containers.Edit(cont))
}

func (h DockerHandler) EditContainer(c echo.Context) error {
	var data services.ContainerData
	c.Bind(&data)

	err := services.Validate.Struct(data)
	if err != nil {
		return RenderError(c, http.StatusBadRequest, err)
	}

	err = services.EditContainer(c.Param("id"), data)
	if err != nil {
		return RenderDockerError(c, err)
	}
	return h.GetContainers(c)
}
