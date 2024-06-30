package handlers

import (
	"net/http"

	"github.com/docker/docker/client"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, err)
	} else if client.IsErrNotFound(err) {
		c.JSON(http.StatusNotFound, err)
	} else if ve, ok := err.(validator.ValidationErrors); ok {
		c.JSON(http.StatusUnprocessableEntity, ve.Error())
	} else {
		c.Logger().Error(err)
		c.String(http.StatusInternalServerError, "Internal server errror")
	}
}
