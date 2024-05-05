package handlers

import (
	"strconv"

	"github.com/RadeJR/containerama/components"
	"github.com/a-h/templ"
	"github.com/docker/docker/api/types"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Locals map[string]interface{}
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func PaginateContainers(cont []types.Container, page int, size int) []types.Container {
	count := len(cont)
	lower := (page - 1) * size
	upper := page * size
	if upper > count {
		upper = count
	}
	return cont[lower:upper]
}

func RenderError(c echo.Context, statusCode int, err error) error {
	c.Response().Header().Set("HX-Retarget", "#popup")
	return Render(c, statusCode, components.ErrorPopup(err))
}

func GetPaginationInfo(c echo.Context) (int, int, error) {
	// PARSING QueryParam
	pageString := c.QueryParam("page")
	var pageNum int
	if pageString != "" {
		var err error
		pageNum, err = strconv.Atoi(pageString)
		if err != nil {
			return -1, -1, err
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
			return -1, -1, err
		}
	} else {
		sizeOfPageNum = 10
	}
	return pageNum, sizeOfPageNum, nil

}
