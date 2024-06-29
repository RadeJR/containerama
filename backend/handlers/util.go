package handlers

import (
	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, err)
	} else if client.IsErrNotFound(err) {
		c.JSON(http.StatusNotFound, err)
	} else {
		c.Logger().Error(err)
		c.String(http.StatusInternalServerError, "Internal server errror")
	}
}
