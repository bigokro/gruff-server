package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func Debates(c echo.Context) error {
	return c.String(http.StatusOK, "Debates")
}
