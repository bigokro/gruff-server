package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func DebateOpinions(c echo.Context) error {
	return c.String(http.StatusOK, "DebateOpinions")
}
