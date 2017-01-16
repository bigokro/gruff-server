package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func ClaimOpinions(c echo.Context) error {
	return c.String(http.StatusOK, "ClaimOpinions")
}
