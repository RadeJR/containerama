package handlers

import (
	"github.com/RadeJR/containerama/components"
	"github.com/RadeJR/containerama/components/networks"
	"github.com/RadeJR/containerama/services"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var NetworkHeaders = []string{"ID", "Driver", "Name", "Created at"}
type NetworkHandler struct{}

func (nh NetworkHandler) ShowNetworks(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	role := sess.Values["role"].(string)

	page, size, err := GetPaginationInfo(c)
	if err != nil {
		return err
	}
	nr, err := services.GetNetworks()
	if err != nil {
		return err
	}
	paginatedNR := PaginateNetworks(nr, page, size)
	td := components.TableData{
		Rows: make([]components.RowData, len(paginatedNR)),
	}

	td.Headers = NetworkHeaders
	for k,v := range paginatedNR {
		td.Rows[k] = services.NewNetworkRowData(v)
	}

	if c.Request().Header.Get("HX-Request") != "true" {
		return Render(c, 200, networks.PageFull(td, page, size, len(nr), role))
	} else {
		return Render(c, 200, networks.Page(td, page, size, len(nr), role))
	}
}
