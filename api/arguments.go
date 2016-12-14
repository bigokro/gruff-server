package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func Arguments(c echo.Context) error {
	return c.String(http.StatusOK, "Arguments")
}
