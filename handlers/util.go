package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Locals map[string]interface{}
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
