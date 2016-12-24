package api

import (
	"github.com/bigokro/gruff-server/gruff"
	"github.com/labstack/echo"
	"net/http"
)

func Debates(c echo.Context) error {
	return c.String(http.StatusOK, "Debates")
}

func (ctx *Context) GetDebate(c echo.Context) error {
	id := c.Param("id")

	db := ctx.Database
	debate := gruff.Debate{}

	err := db.Preload("ProTruth", "parent_id = ?", id).Preload("ConTruth", "parent_id = ?", id).Preload("Links").Preload("Contexts").Preload("Values").Preload("Tags").Where("id = ?", id).Find(&debate).Error
	if err != nil {
		return gruff.NewServerError(err.Error())
	}

	return c.JSON(http.StatusOK, debate)
}
