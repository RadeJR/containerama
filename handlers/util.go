package handlers

import (
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
