package handlers

import (
	"log/slog"
	"net/http"

	"github.com/RadeJR/containerama/components"
	"github.com/a-h/templ"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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

func RenderDockerError(c echo.Context, err error) error {
	if client.IsErrConnectionFailed(err) {
		slog.Error("Error connecting to docker", "error", err.Error())
		return c.String(http.StatusInternalServerError, "Internal server error")
	}
	if client.IsErrNotFound(err) {
		return RenderError(c, http.StatusNotFound, err)
	}
	return RenderError(c, http.StatusBadRequest, err)
}
