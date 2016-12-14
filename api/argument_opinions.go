package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func ArgumentOpinions(c echo.Context) error {
	return c.String(http.StatusOK, "ArgumentOpinions")
}
