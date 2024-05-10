package handlers

import (
	"github.com/RadeJR/containerama/components"
	compstacks "github.com/RadeJR/containerama/components/stacks"
	"github.com/RadeJR/containerama/models"
	"github.com/RadeJR/containerama/services"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var stackHeaders = []string{"ID", "Name", "Created At"}
type StackHandler struct{}

func (sh StackHandler) GetStacks(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	role := sess.Values["role"].(string)
	id := sess.Values["id"].(int)

	page, size, err := GetPaginationInfo(c)
	if err != nil {
		return err
	}

	var stacks []models.Stack
	stacks, err = services.GetStacks(id, role, page, size)
	if err != nil {
		return err
	}

	td := components.TableData{}
	td.Rows = make([]components.RowData, len(stacks))
	td.Headers = stackHeaders
	for k,v := range stacks {
		td.Rows[k] = services.NewStackRowData(v)
	}

	if c.Request().Header.Get("HX-Request") != "true" {
		return Render(c, 200, compstacks.PageFull(td, page, size, len(stacks), role))
	} else {
		return Render(c, 200, compstacks.Page(td, page, size, len(stacks), role))
	}
}
