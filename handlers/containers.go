package handlers

import (
	"github.com/RadeJR/containerama/components"
	"github.com/RadeJR/containerama/components/containers"
	"github.com/RadeJR/containerama/services"
	"github.com/labstack/echo/v4"
)

type DockerHandler struct{}

func (h DockerHandler) GetContainers(c echo.Context) error {
	page, size, err := GetPaginationInfo(c)
	if err != nil {
		return err
	}
	cont, err := services.GetContainers()
	if err != nil {
		return err
	}
	paginatedCont := PaginateContainers(cont, page, size)

	tableData := components.TableData{
		Rows: make([]components.RowData, len(paginatedCont)),
	}
	tableData.Headers = containers.Headers
	for k, v := range paginatedCont {
		tableData.Rows[k] = services.NewRowData(v)
	}

	role := c.(CustomContext).Locals["role"].(string)
	if c.Request().Header.Get("HX-Request") != "true" {
		return Render(c, 200, containers.PageFull(tableData, page, size, len(cont), role))
	} else {
		return Render(c, 200, containers.Page(tableData, page, size, len(cont), role))
	}
}

func (h DockerHandler) StopContainer(c echo.Context) error {
	err := services.StopContainer(c.Param("id"))
	if err != nil {
		return err
	}
	return h.GetContainers(c)
}

func (h DockerHandler) CreateContainerPage(c echo.Context) error {
	return Render(c, 200, containers.Create())
}

func (h DockerHandler) CreateContainer(c echo.Context) error {
	var data services.ContainerData
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	err = services.Validate.Struct(data)
	if err != nil {
		return err
	}

	err = services.CreateContainer(data)
	if err != nil {
		return err
	}
	return h.GetContainers(c)
}

func (h DockerHandler) StartContainer(c echo.Context) error {
	err := services.StartContainer(c.Param("id"))
	if err != nil {
		return err
	}
	return h.GetContainers(c)
}

func (h DockerHandler) RestartContainer(c echo.Context) error {
	err := services.RestartContainer(c.Param("id"))
	if err != nil {
		return err
	}
	return h.GetContainers(c)
}

func (h DockerHandler) RemoveContainer(c echo.Context) error {
	err := services.RemoveContainer(c.Param("id"), false)
	if err != nil {
		return err
	}
	return h.GetContainers(c)
}

func (h DockerHandler) ShowContainer(c echo.Context) error {
	cont, err := services.GetContainer(c.Param("id"))
	if err != nil {
		return err
	}
	return Render(c, 200, containers.One(cont))
}

func (h DockerHandler) EditContainerPage(c echo.Context) error {
	cont, err := services.GetContainer(c.Param("id"))
	if err != nil {
		return err
	}
	return Render(c, 200, containers.Edit(cont))
}

func (h DockerHandler) EditContainer(c echo.Context) error {
	var data services.ContainerData
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	err = services.Validate.Struct(data)
	if err != nil {
		return err
	}

	err = services.EditContainer(c.Param("id"), data)
	if err != nil {
		return err
	}
	return h.GetContainers(c)
}
